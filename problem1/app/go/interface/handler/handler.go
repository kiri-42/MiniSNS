package handler

import (
	"net/http"
	"problem1/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	Root() echo.HandlerFunc
	GetUser() echo.HandlerFunc
	GetUserList() echo.HandlerFunc
	GetUserListPaging() echo.HandlerFunc
	GetFriendList() echo.HandlerFunc
	GetFriendOfFriendList() echo.HandlerFunc
	GetFriendOfFriendListPaging() echo.HandlerFunc
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

// NewUserHandler はuserHandlerのコンストラクタです。
func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{userUsecase: userUsecase}
}

type resUser struct {
	UserID int     `json:"user_id"`
	Name   string  `json:"name"`
}

// Root は"/"のhttpハンドラです。
// "mini sns"を返します
func (uh *userHandler) Root() echo.HandlerFunc  {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "mini sns")
	}
}

// GetUser は"/get_user/:user_id"のhttpハンドラです。
// user_idをもとにUserをjson形式で返します。
func (uh *userHandler) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		uID, err := strconv.Atoi(c.Param("user_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		u, err := uh.userUsecase.GetUser(uID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, getResUser(u))
	}
}

// GetUserList は"/get_user_list"のhttpハンドラです。
// User listをjson形式で返します。
func (uh *userHandler) GetUserList() echo.HandlerFunc {
	return func(c echo.Context) error {
		uList, err := uh.userUsecase.GetUserList()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, getResUserList(uList))
	}
}

// GetUserListPaging は"/get_user_list_page"のhttpハンドラです。
// limitとpageをもとにUser listをjson形式で返します。
func (uh *userHandler) GetUserListPaging() echo.HandlerFunc {
	return func(c echo.Context) error {
		limit, err := strconv.Atoi(c.Param("limit"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		page, err := strconv.Atoi(c.Param("page"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		uList, err := uh.userUsecase.GetUserListPaging(limit, page)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, getResUserList(uList))
	}
}

// GetFriendList は"/get_friend_list/:user_id"のhttpハンドラです。
// user_idをもとにfriend listをjson形式で返します。
func (uh *userHandler) GetFriendList() echo.HandlerFunc {
	return func(c echo.Context) error {
		uID, err := strconv.Atoi(c.Param("user_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		fList, err := uh.userUsecase.GetFriendList(uID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, getResUserList(fList))
	}
}

// GetFriendOfFriendList は"/get_friend_of_friend_list/:user_id"のhttpハンドラです。
// user_idをもとにfriendのfriend listをjson形式で返します。
func (uh *userHandler) GetFriendOfFriendList() echo.HandlerFunc {
	return func(c echo.Context) error {
		uID, err := strconv.Atoi(c.Param("user_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		ffList, err := uh.userUsecase.GetFriendOfFriendList(uID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, getResUserList(ffList))
	}
}

// GetFriendOfFriendListPaging は"/get_friend_of_friend_list_paging/:user_id"のhttpハンドラです。
// user_idとlimit,pageをもとにfriendのfriend listをjson形式で返します。
func (uh *userHandler) GetFriendOfFriendListPaging() echo.HandlerFunc {
	return func(c echo.Context) error {
		uID, err := strconv.Atoi(c.Param("user_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		limit, err := strconv.Atoi(c.Param("limit"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		page, err := strconv.Atoi(c.Param("page"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		ffList, err := uh.userUsecase.GetFriendOfFriendListPaging(uID, limit, page)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, getResUserList(ffList))
	}
}

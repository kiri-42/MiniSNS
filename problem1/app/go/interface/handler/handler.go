package handler

import (
	"net/http"
	"problem1/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandlerI interface {
	Root() echo.HandlerFunc
	GetUser() echo.HandlerFunc
	GetUserList() echo.HandlerFunc
	GetUserListPaging() echo.HandlerFunc
	GetFriendList() echo.HandlerFunc
	GetFriendOfFriendList() echo.HandlerFunc
	GetFriendOfFriendListPaging() echo.HandlerFunc
}

type userHandlerS struct {
	userUsecase usecase.UserUsecaseI
}

// NewUserHandler はuserHandlerのコンストラクタです。
func NewUserHandler(userUsecase usecase.UserUsecaseI) UserHandlerI {
	return &userHandlerS{userUsecase: userUsecase}
}

type resUser struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
}

// Root は"/"のhttpハンドラです。
// "mini sns"を返します
func (uh *userHandlerS) Root() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "mini sns")
	}
}

// GetUser は"/get_user/:user_id"のhttpハンドラです。
// user_idをもとにUserをjson形式で返します。
func (uh *userHandlerS) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Param("user_id") == "" {
			return newHTTPError(http.StatusInternalServerError, "limit is user_id")
		}
		uID, err := strconv.Atoi(c.Param("user_id"))
		if err != nil {
			return newHTTPError(http.StatusInternalServerError, err.Error())
		}

		u, err := uh.userUsecase.GetUser(uID)
		if err != nil {
			return newHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, getResUser(u))
	}
}

// GetUserList は"/get_user_list"のhttpハンドラです。
// User listをjson形式で返します。
func (uh *userHandlerS) GetUserList() echo.HandlerFunc {
	return func(c echo.Context) error {
		uList, err := uh.userUsecase.GetUserList()
		if err != nil {
			return newHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, getResUserList(uList))
	}
}

// GetUserListPaging は"/get_user_list_page"のhttpハンドラです。
// limitとpageをもとにUser listをjson形式で返します。
func (uh *userHandlerS) GetUserListPaging() echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Param("limit") == "" {
			return newHTTPError(http.StatusInternalServerError, "limit is empty")
		}
		limit, err := strconv.Atoi(c.Param("limit"))
		if err != nil {
			return newHTTPError(http.StatusInternalServerError, err.Error())
		}

		if c.Param("page") == "" {
			return newHTTPError(http.StatusInternalServerError, "page is empty")
		}
		page, err := strconv.Atoi(c.Param("page"))
		if err != nil {
			return newHTTPError(http.StatusInternalServerError, err.Error())
		}

		uList, err := uh.userUsecase.GetUserListPaging(limit, page)
		if err != nil {
			return newHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, getResUserList(uList))
	}
}

// GetFriendList は"/get_friend_list/:user_id"のhttpハンドラです。
// user_idをもとにfriend listをjson形式で返します。
func (uh *userHandlerS) GetFriendList() echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Param("user_id") == "" {
			return newHTTPError(http.StatusInternalServerError, "user_id is empty")
		}
		uID, err := strconv.Atoi(c.Param("user_id"))
		if err != nil {
			return newHTTPError(http.StatusInternalServerError, err.Error())
		}

		fList, err := uh.userUsecase.GetFriendList(uID)
		if err != nil {
			return newHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, getResUserList(fList))
	}
}

// GetFriendOfFriendList は"/get_friend_of_friend_list/:user_id"のhttpハンドラです。
// user_idをもとにfriendのfriend listをjson形式で返します。
func (uh *userHandlerS) GetFriendOfFriendList() echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Param("user_id") == "" {
			return newHTTPError(http.StatusInternalServerError, "user_id is empty")
		}
		uID, err := strconv.Atoi(c.Param("user_id"))
		if err != nil {
			return newHTTPError(http.StatusInternalServerError, err.Error())
		}

		ffList, err := uh.userUsecase.GetFriendOfFriendList(uID)
		if err != nil {
			return newHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, getResUserList(ffList))
	}
}

// GetFriendOfFriendListPaging は"/get_friend_of_friend_list_paging/:user_id"のhttpハンドラです。
// user_idとlimit,pageをもとにfriendのfriend listをjson形式で返します。
func (uh *userHandlerS) GetFriendOfFriendListPaging() echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Param("user_id") == "" {
			return newHTTPError(http.StatusInternalServerError, "user_id is empty")
		}
		uID, err := strconv.Atoi(c.Param("user_id"))
		if err != nil {
			return newHTTPError(http.StatusInternalServerError, err.Error())
		}

		if c.Param("limit") == "" {
			return newHTTPError(http.StatusInternalServerError, "limit is empty")
		}
		limit, err := strconv.Atoi(c.Param("limit"))
		if err != nil {
			return newHTTPError(http.StatusInternalServerError, err.Error())
		}

		if c.Param("page") == "" {
			return newHTTPError(http.StatusInternalServerError, "page is empty")
		}
		page, err := strconv.Atoi(c.Param("page"))
		if err != nil {
			return newHTTPError(http.StatusInternalServerError, err.Error())
		}

		ffList, err := uh.userUsecase.GetFriendOfFriendListPaging(uID, limit, page)
		if err != nil {
			return newHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, getResUserList(ffList))
	}
}

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
	GetFriendList() echo.HandlerFunc
	GetFriendOfFriendList() echo.HandlerFunc
	GetFriendOfFriendListPaging() echo.HandlerFunc
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{userUsecase: userUsecase}
}

type resUser struct {
	UserID int     `json:"user_id"`
	Name   string  `json:"name"`
}

func (uh *userHandler) Root() echo.HandlerFunc  {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "mini sns")
	}
}

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
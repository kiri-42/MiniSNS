package handler

import (
	"net/http"
	"problem1/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	Root() echo.HandlerFunc
	Get() echo.HandlerFunc
	GetFriendList() echo.HandlerFunc
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

func (uh *userHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		foundUser, err := uh.userUsecase.FindByID(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		res := resUser {
			UserID: foundUser.UserID,
			Name:   foundUser.Name,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (uh *userHandler) GetFriendList() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		foundFriendList, err := uh.userUsecase.FindFriendsByID(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		res := make([]resUser, 0)
		for _, v := range foundFriendList {
			friend := resUser {
				UserID: v.UserID,
				Name:   v.Name,
			}
			res = append(res, friend)
		}

		return c.JSON(http.StatusOK, res)
	}
}

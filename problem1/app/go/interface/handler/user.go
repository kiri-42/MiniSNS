package handler

import (
	"net/http"
	"problem1/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	Get() echo.HandlerFunc
	Root() echo.HandlerFunc
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



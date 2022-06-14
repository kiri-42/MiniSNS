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

		id, err := uh.userUsecase.FindIDByUserID(uID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		u, err := uh.userUsecase.FindByID(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		res := resUser {
			UserID: u.UserID,
			Name:   u.Name,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (uh *userHandler) GetFriendList() echo.HandlerFunc {
	return func(c echo.Context) error {
		uID, err := strconv.Atoi(c.Param("user_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		id, err := uh.userUsecase.FindIDByUserID(uID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		fList, err := uh.userUsecase.FindFriendListExceptBlock(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		fList = uh.userUsecase.GetUniqueList(fList)

		res := make([]resUser, 0)
		for _, v := range fList {
			friend := resUser {
				UserID: v.UserID,
				Name:   v.Name,
			}
			res = append(res, friend)
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (uh *userHandler) GetFriendOfFriendList() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := strconv.Atoi(c.Param("user_id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		id, err := uh.userUsecase.FindIDByUserID(userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		fList, err := uh.userUsecase.FindFriendListExceptBlock(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		ffList, err := uh.userUsecase.FindFriendOfFriendListExcept1HopFriend(fList)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		ffList = uh.userUsecase.GetUniqueList(ffList)

		res := make([]resUser, 0)
		for _, f := range ffList {
			user := resUser {
				UserID: f.UserID,
				Name:   f.Name,
			}
			res = append(res, user)
		}

		return c.JSON(http.StatusOK, res)
	}
}

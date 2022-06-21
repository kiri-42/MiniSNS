package handler

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"problem1/infrastructure"
	"problem1/usecase"
)

func NewRouter(db *sql.DB) *echo.Echo {
	userRepository := infrastructure.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := NewUserHandler(userUsecase)

	e := echo.New()
	e.HTTPErrorHandler = HttpErrorHandler
	e.Use(middleware.Recover()) // httpハンドラ内でpanicしてもサーバーが落ちないようにする

	Routing(e, userHandler)

	return e
}

func Routing(e *echo.Echo, userHandler UserHandlerI) {
	e.GET("/", userHandler.Root())
	e.GET("/get_user/:user_id", userHandler.GetUser())
	e.GET("/get_user_list", userHandler.GetUserList())
	e.GET("/get_user_list_paging/:limit/:page", userHandler.GetUserListPaging())
	e.GET("/get_friend_list/:user_id", userHandler.GetFriendList())
	e.GET("/get_friend_of_friend_list/:user_id", userHandler.GetFriendOfFriendList())
	e.GET("/get_friend_of_friend_list_paging/:user_id/:limit/:page", userHandler.GetFriendOfFriendListPaging())
}

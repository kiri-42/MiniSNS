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
	e.Use(middleware.Logger())  // httpリクエストのロクを出力

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

// package handler_test

// import (
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	// "strings"
// 	"testing"

// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/echo/v4/middleware"

// 	"problem1/configs"
// 	"problem1/infrastructure"
// 	"problem1/interface/handler"
// 	"problem1/usecase"
// )

// func TestRoot(t *testing.T) {
// 	db, err := configs.GetDB()
// 	if err != nil {
// 		fmt.Fprintln(os.Stdout, err.Error())
// 		return
// 	}
// 	defer db.Close()

// 	// userRepository := infrastructure.NewUserRepository(db)
// 	// userUsecase := usecase.NewUserUsecase(userRepository)
// 	// userHandler := handler.NewUserHandler(userUsecase)

// 	// e := echo.New()
// 	// e.HTTPErrorHandler = handler.HttpErrorHandler
// 	// e.Use(middleware.Recover()) // httpハンドラ内でpanicしてもサーバーが落ちないようにする
// 	// e.Use(middleware.Logger())  // httpリクエストのロクを出力

// 	// handler.Routing(e, userHandler)

// 	e := main.NewRouter()

// 	req := httptest.NewRequest("GET", "/", nil)
// 	rec := httptest.NewRecorder()

// 	e.ServeHTTP(rec, req)

// 	if http.StatusOK == rec.Code {
// 		println("OK")
// 	} else {
// 		println("NG")
// 	}
// }


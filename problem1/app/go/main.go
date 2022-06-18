package main

import (
	"fmt"
	"os"
	"problem1/configs"
	"problem1/infrastructure"
	"problem1/interface/handler"
	"problem1/usecase"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	conf := configs.Get()

	db, err := configs.GetDB()
	if err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
		return
	}
	defer db.Close()

	userRepository :=  infrastructure.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	e := echo.New()
	e.Use(middleware.Recover()) // httpハンドラ内でpanicしてもサーバーが落ちないようにする
	e.Use(middleware.Logger())  // httpリクエストのロクを出力

	handler.Routing(e, userHandler)
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(conf.Server.Port)))
}

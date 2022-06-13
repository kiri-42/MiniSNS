package main

import (
	"problem1/configs"
	"problem1/interface/handler"
	"problem1/infrastructure"
	"problem1/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	conf := configs.Get()

	db, err := configs.GetDB()
	if err != nil {
		println(err.Error())
	}
	defer db.Close()

	userRepository :=  infrastructure.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	e := echo.New()
	handler.Routing(e, userHandler)
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(conf.Server.Port)))
}

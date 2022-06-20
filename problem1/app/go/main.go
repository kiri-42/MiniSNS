package main

import (
	_ "database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/labstack/echo/v4"

	"problem1/configs"
	"problem1/interface/handler"
)

func main() {
	conf := configs.Get()

	db, err := configs.GetDB()
	if err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
		return
	}
	defer db.Close()

	e := handler.NewRouter(db)
	if err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
		return
	}

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(conf.Server.Port)))
}

package handler

import "github.com/labstack/echo/v4"

func Routing(e *echo.Echo, userHandler UserHandler) {
	e.GET("/get_user/:id", userHandler.Get())
	e.GET("/", userHandler.Root())
}

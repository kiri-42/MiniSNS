package handler

import "github.com/labstack/echo/v4"

func Routing(e *echo.Echo, userHandler UserHandler) {
	e.GET("/", userHandler.Root())
	e.GET("/get_user/:id", userHandler.Get())
	e.GET("/get_friend_list/:id", userHandler.GetFriendList())
}

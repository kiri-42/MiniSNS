package handler

import "github.com/labstack/echo/v4"

func Routing(e *echo.Echo, userHandler UserHandler) {
	e.GET("/", userHandler.Root())
	e.GET("/get_user/:user_id", userHandler.GetUser())
	e.GET("/get_user_list", userHandler.GetUserList())
	e.GET("/get_user_list_paging/:limit/:page", userHandler.GetUserListPaging())
	e.GET("/get_friend_list/:user_id", userHandler.GetFriendList())
	e.GET("/get_friend_of_friend_list/:user_id", userHandler.GetFriendOfFriendList())
	e.GET("/get_friend_of_friend_list_paging/:user_id/:limit/:page", userHandler.GetFriendOfFriendListPaging())
}

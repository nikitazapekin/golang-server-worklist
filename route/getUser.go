package router

import (
	"server/controller"
	"github.com/labstack/echo/v4"
)
func SetGetUsers(e *echo.Echo) {
	e.GET("/worklist.com/getUser", controller.GetUsers)
}


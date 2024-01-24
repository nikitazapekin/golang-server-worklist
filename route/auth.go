package router

import (
	"server/controller"
	"github.com/labstack/echo/v4"
)

func SetAuth(e *echo.Echo) {
	e.GET("/auth/register", controller.Register)
}


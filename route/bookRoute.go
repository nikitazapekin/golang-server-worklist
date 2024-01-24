package router

import (
	"server/controller"
	"github.com/labstack/echo/v4"
)

func SetBookRoutes(e *echo.Echo) {
	e.GET("/books", controller.GetBooks)
}


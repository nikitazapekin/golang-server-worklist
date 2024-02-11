package router

import (
	"server/controller"
	"github.com/labstack/echo/v4"
)

func SetFilteredRoutes(e *echo.Echo) {
	e.GET("/worklist.com/filterOffers", controller.GetBooks)
}


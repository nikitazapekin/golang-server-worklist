package router

import (
	"server/controller"
	"github.com/labstack/echo/v4"
)

func SetFilteredRoutes(e *echo.Echo) {
	e.POST("/worklist.com/filterOffers", controller.GetFilteredOffers)
}


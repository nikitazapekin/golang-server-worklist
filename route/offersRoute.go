package router

import (
	"server/controller"
	"github.com/labstack/echo/v4"
)

func SetOffersRoutes(e *echo.Echo) {
	e.GET("/worklist.com/getOffers", controller.GetOffers)
}


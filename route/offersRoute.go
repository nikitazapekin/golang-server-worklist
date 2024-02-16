package router

import (
	"server/controller"
	"github.com/labstack/echo/v4"
)

func SetOffersRoutes(e *echo.Echo) {
	//e.GET("/worklist.com/getOffers", controller.GetOffers)
	e.POST("/worklist.com/getOffers", controller.GetOffers)
	e.GET("/worklist.com/getAmountOfOffers", controller.GetAmountOfOffers)
	e.GET("/worklist.com/getOffer", controller.GetOffer)
	e.GET("/worklist.com/getYourOffers", controller.GetYourOffers)
}


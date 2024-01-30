package router

import (
	"server/controller"
	"github.com/labstack/echo/v4"
)

func SetPersonalInformation(e *echo.Echo) {
	//e.GET("/worklist.com/getPersonalInformation/:token", controller.GetPersonalInformation)
	e.GET("/worklist.com/getPersonalInformation", controller.GetPersonalInformation)
}


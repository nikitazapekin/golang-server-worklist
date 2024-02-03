package router

import (
	"server/controller"
	"github.com/labstack/echo/v4"
)

func SetVacancy(e *echo.Echo) {
	e.POST("/worklist.com/createVacancy", controller.CreateVacancy)
}


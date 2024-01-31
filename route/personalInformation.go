
package router

import (
	"server/controller"
	"github.com/labstack/echo/v4"
)
func SetPersonalInformation(e *echo.Echo) {
	e.POST("/worklist.com/getPersonalInformation/editPersonalData", controller.EditPersonalInformation)
	e.GET("/worklist.com/getPersonalInformation", controller.GetPersonalInformation)
}


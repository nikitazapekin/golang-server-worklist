
package router
import (
	"server/controller"
	"github.com/labstack/echo/v4"
	//"github.com/gin-gonic/gin"
)
/*
func SetAvatarHandler(c echo.Context) error {
    // Create a new echo.Context using the request and response from gin.Context
    eContext := echo.New().NewContext(c.Request(), c.Response())

    // Call the original controller.SetAvatar with the new echo.Context
    controller.SetAvatar(eContext)

    return nil
} */
func SetPersonalInformation(e *echo.Echo) {
	e.POST("/worklist.com/getPersonalInformation/editPersonalData", controller.EditPersonalInformation)
	e.GET("/worklist.com/getPersonalInformation", controller.GetPersonalInformation)
//	e.Static("/worklist.com/getPersonalInformation/setAvatar/", "static")
//e.Static("/worklist.com/", "static")
//e.Static("/", "static")
//e.Use(middleware.Static("/static"))
//e.POST("/worklist.com/getPersonalInformation/setAvatar",SetAvatarHandler)
	e.POST("/worklist.com/getPersonalInformation/setAvatar", controller.SetAvatar)
}


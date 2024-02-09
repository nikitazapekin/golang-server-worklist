package router

import (
	"server/controller"
	"github.com/labstack/echo/v4"
)

func SetGetPhotos(e *echo.Echo) {

    //e.GET("/worklist.com/image", controller.Imgg)
	e.GET("/worklist.com/image", controller.ImgEcho)
}
package controller

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
)


func GetUsers(c echo.Context) error {
	username:= c.QueryParam("username")
fmt.Println(username)
	return c.JSON(http.StatusOK, "kd")
}
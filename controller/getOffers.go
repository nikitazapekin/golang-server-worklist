package controller

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	m "server/db"
)
func GetOffers(c echo.Context) error {
fmt.Println("HELOOOOOOOOOOOOOOOOOOO")
elem, err:=m.GetAllVacancyData()
fmt.Println(err)
fmt.Println(elem)
if(err!=nil) {
	return c.JSON(http.StatusBadRequest, "{message: error}")
}
	return c.JSON(http.StatusOK, elem)
}
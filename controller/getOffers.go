package controller

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	"strconv"
	m "server/db"
)

func GetOffers(c echo.Context) error {

	limitStr := c.QueryParam("limit")
    pageStr := c.QueryParam("page")
	
fmt.Println("HELOOOOOOOOOOOOOOOOOOO")
limit, err := strconv.Atoi(limitStr)
if err != nil {
	return c.JSON(http.StatusBadRequest, "Invalid limit parameter")
}
page, err := strconv.Atoi(pageStr)
if err != nil {
	return c.JSON(http.StatusBadRequest, "Invalid page parameter")
}
elem, err := m.GetAllVacancyData(limit, page)
//elem, err:=m.GetAllVacancyData()
fmt.Println(err)
fmt.Println(elem)
if(err!=nil) {
	return c.JSON(http.StatusBadRequest, "{message: error}")
}
	return c.JSON(http.StatusOK, elem)
} 
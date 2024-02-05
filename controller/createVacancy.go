package controller

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	"encoding/json"
	e "server/middleware"
) 
type CreateVacancyParams struct {
	Title string `json:"title"`
	Describtion string `json:"describtion"`
	Skills []string `json:"skills"`
	WorkingPerDay string `json:"workingPerDay"`
	Location string `json:"location"`
	Salary string `json:"salary"`
}
func CreateVacancy(c echo.Context) error {
	token := c.QueryParam("token")
	fmt.Println("Received token:", token)
	decodedToken, errToken := e.Decode(token,  "key")
fmt.Println("Decoded")
fmt.Println(decodedToken)
fmt.Println(errToken)
var createVacancyParams  CreateVacancyParams
	err := json.NewDecoder(c.Request().Body).Decode(&createVacancyParams)
	fmt.Println(err)
return c.JSON(http.StatusOK, "{message: error}")
//	return c.JSON(http.StatusOK, u)
}
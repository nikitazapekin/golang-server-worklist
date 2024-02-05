package controller

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	"encoding/json"
	e "server/middleware"
	"server/db"
) 
type CreateVacancyParams struct {
	Title string `json:"title"`
	Describtion string `json:"describtion"`
	Skills []string `json:"skills"`
	WorkingPerDay string `json:"workingPerDay"`
	Location string `json:"location"`
	Salary string `json:"salary"`
}
//var arrayOfLastImgs String[]=[]

func CreateVacancy(c echo.Context) error {
	token := c.QueryParam("token")
	fmt.Println("Received token:", token)
	decodedToken, errToken := e.Decode(token,  "key")
fmt.Println("Decoded")
fmt.Println(decodedToken)
fmt.Println(errToken)
var createVacancyParams  CreateVacancyParams
	err := json.NewDecoder(c.Request().Body).Decode(&createVacancyParams)
	err=db.InsertDataIntoOffers(c.Response(), createVacancyParams.Title, createVacancyParams.Describtion, createVacancyParams.Skills, createVacancyParams.WorkingPerDay, createVacancyParams.Location, createVacancyParams.Salary )
	fmt.Println(createVacancyParams.Title)
	fmt.Println(err)
return c.JSON(http.StatusOK, "{message: error}")
//	return c.JSON(http.StatusOK, u)
}
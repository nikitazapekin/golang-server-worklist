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
	Token string `json:"token"`
	ArrayOfPictures []string `json:"arrayOfPictures"`
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
	fmt.Println("ARAYYYYYYYYYYYYYYYYYYYYYY ")
	fmt.Println(createVacancyParams.ArrayOfPictures)
	createVacancyParams.Token = token
	id, err:=db.InsertDataIntoOffers(c.Response(), createVacancyParams.Title, createVacancyParams.Describtion, createVacancyParams.Skills, createVacancyParams.WorkingPerDay, createVacancyParams.Location, createVacancyParams.Salary, createVacancyParams.Token, createVacancyParams.ArrayOfPictures )
	fmt.Println(createVacancyParams.Title)
	fmt.Println(err)
	fmt.Println("CURRRRRRRRRRRRRRRRRRRRRRRRRRRRRREEEEEEEEEEEEEEEEENT IDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDD")
	fmt.Println(id)
	user, err := db.FindUserByUsername(decodedToken.Username)
	db.UpdateUserIds(user, decodedToken.Username, id)
return c.JSON(http.StatusOK, "{message: error}")
}
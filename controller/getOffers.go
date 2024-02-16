package controller

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	"strconv"
	"encoding/json"
	m "server/db"
)
type SearchVacancyParams struct {
	Title         string   `json:"title"`
	Skills        []string `json:"skills"`
	WorkingPerDay string   `json:"workingPerDay"`
	Location      string   `json:"location"`
	SalaryFrom    string   `json:"from"`
	SalaryTo      string   `json:"to"`
}
func GetOffers(c echo.Context) error {
	var searchParams SearchVacancyParams
	err := json.NewDecoder(c.Request().Body).Decode(&searchParams)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "error"})
	}

	fmt.Println("===================================================================================================================================================================================")
	fmt.Println("PROOOOOOOOOOOOOOOOOOOOPS")
	fmt.Println(searchParams) 
fmt.Println("TITLE"+searchParams.Title)
fmt.Println("SKILLS")
fmt.Println(searchParams.Skills)
fmt.Println("WORKING PER DAY"+searchParams.WorkingPerDay)
fmt.Println("Location" +searchParams.Location)
fmt.Println("FROM" +searchParams.SalaryFrom)
fmt.Println(("TO" +searchParams.SalaryTo)) 

fmt.Println("===================================================================================================================================================================================")
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
if searchParams.Title=="" && len(searchParams.Skills)==0 && searchParams.WorkingPerDay=="0" && searchParams.Location=="" && searchParams.SalaryFrom == "" && searchParams.SalaryTo==""{
//	if searchParams.Title==""  && searchParams.SalaryTo==""{
	elem, err := m.GetAllVacancyData(limit, page)
	//fmt.Println(err)
//fmt.Println(elem)
if(err!=nil) {
	return c.JSON(http.StatusBadRequest, "{message: error}")
}
return c.JSON(http.StatusOK, elem) 
} else {
	fmt.Println("HANDLE SEARCHhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh")
	title:= searchParams.Title
	skills := searchParams.Skills
	workingPerDay := searchParams.WorkingPerDay
	location := searchParams.Location
	from:=searchParams.SalaryFrom
	to:=searchParams.SalaryTo
 
	elem, err := m.GetFilteredPaginationData(limit, page, title, skills, workingPerDay, location, from, to)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "error"})
	}

	fmt.Println("FINALLY ")
	fmt.Println(elem)
//	return nil

return c.JSON(http.StatusOK, elem)
}
} 
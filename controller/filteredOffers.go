/*package controller
import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	m "server/db"
)
type SearchVacancyParams struct {
	Title string `json:"title"`
	Skills []string `json:"skills"`
	WorkingPerDay string `json:"workingPerDay"`
	Location string `json:"location"`
	SalaryFrom string `json:"from"`
	SalaryTo string `json:"to"`
}
func GetFilteredOffers(c echo.Context) error {
	var searchParams SearchVacancyParams
	err := json.NewDecoder(c.Request().Body).Decode(&searchParams)
	if(err!=nil) {
		return c.JSON(http.StatusBadRequest, "{error: error}")
	} 
fmt.Println("HELOOOOOOOOOOOOOOOOOOO")
fmt.Println(searchParams)
elems, err := m.GetAllOffers()
fmt.Println(elems)
	return c.JSON(http.StatusOK, "hello")
}
*/


/*
package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
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

func GetFilteredOffers(c echo.Context) error {
	var searchParams SearchVacancyParams
	err := json.NewDecoder(c.Request().Body).Decode(&searchParams)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "error"})
	}

	fmt.Println("HELOOOOOOOOOOOOOOOOOOO")
	fmt.Println(searchParams)

	elems, err := m.GetAllOffers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to fetch offers"})
	}

	filteredOffers := make([]m.VacancyData, 0)
	for _, elem := range elems {
		if containsTitle(elem.Title, searchParams.Title) {
			filteredOffers = append(filteredOffers, elem)
		}
	}

	fmt.Println(filteredOffers)
	return c.JSON(http.StatusOK, filteredOffers)
}

func containsTitle(title, searchTitle string) bool {
	return title == searchTitle
}  */



package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings" 

	"github.com/labstack/echo/v4"
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
func GetFilteredOffers(c echo.Context) error {
	var searchParams SearchVacancyParams
	err := json.NewDecoder(c.Request().Body).Decode(&searchParams)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "error"})
	}
	fmt.Println(searchParams)
	elems, err := m.GetAllOffers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to fetch offers"})
	}
	searchTitle := strings.ToLower(searchParams.Title)
	searchLocation := strings.ToLower(searchParams.Location)
	filteredOffers := make([]m.VacancyData, 0)        // ГЛОБАЛЬНОЕ ХРАНИЛИЩЕ ОТФИЛЬТРОВАННЫХ ПРЕДЛОЖЕНИЙ
	//====================================================================================
	//фильтрация
	for _, elem := range elems {
		isAbleToAddIntoFilterdItems := true
		title := strings.ToLower(elem.Title)
		location := strings.ToLower(elem.Location)
		if len(searchTitle) > 0 {
		if containsTitle(title, searchTitle) {
		//	filteredOffers = append(filteredOffers, elem)
		} else {
			isAbleToAddIntoFilterdItems=false
		}
}
if len(searchLocation) > 0 {
	if containsTitle(location, searchLocation) {
	//	filteredOffers = append(filteredOffers, elem)
	} else {
		isAbleToAddIntoFilterdItems=false
	}
}
 if isAbleToAddIntoFilterdItems==true {
	filteredOffers = append(filteredOffers, elem)
} else {
	//filteredOffers=elems
}
	}
//========================================================================================
	fmt.Println(filteredOffers)
	return c.JSON(http.StatusOK, filteredOffers)
}

func containsTitle(title, searchTitle string) bool {
	//return title == searchTitle
	return strings.Contains(title, searchTitle)
}

 
/*
{
     "title": "test",
     "skills": ["test"],
     "workingPerDay": "test",
     "location": "test",
     "from": "from",
     "to": "to"
}
*/
/*
	var editDataParams EditDataParams
	err := json.NewDecoder(c.Request().Body).Decode(&editDataParams)
	*/
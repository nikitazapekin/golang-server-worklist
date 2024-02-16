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
	//"encoding/json"
	//"fmt"
	"net/http"
	"strings" 
	//"strconv"
	"github.com/labstack/echo/v4"
	//m "server/db"
)
/*
type SearchVacancyParams struct {
	Title         string   `json:"title"`
	Skills        []string `json:"skills"`
	WorkingPerDay string   `json:"workingPerDay"`
	Location      string   `json:"location"`
	SalaryFrom    string   `json:"from"`
	SalaryTo      string   `json:"to"`
} */
func GetFilteredOffers(c echo.Context) error {
	//var searchParams SearchVacancyParams

	/*
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
	searchWorkingPerDay := strings.ToLower(searchParams.WorkingPerDay)
searchSkills := searchParams.Skills
searchSkillsLower := make([]string, len(searchSkills))
for i, v := range searchSkills {
searchSkillsLower[i] = strings.ToLower(v)
}

searchFrom, err := strconv.Atoi(searchParams.SalaryFrom)
	searchTo, err := strconv.Atoi(searchParams.SalaryTo)
	salaryBuffer :=0

		if searchTo < searchFrom {
salaryBuffer = searchFrom
searchFrom=searchTo
searchTo=salaryBuffer
	}
	filteredOffers := make([]m.VacancyData, 0)        // ГЛОБАЛЬНОЕ ХРАНИЛИЩЕ ОТФИЛЬТРОВАННЫХ ПРЕДЛОЖЕНИЙ
	//====================================================================================
	//фильтрация
	for _, elem := range elems {
		isAbleToAddIntoFilterdItems := true
		title := strings.ToLower(elem.Title)
		location := strings.ToLower(elem.Location)
		workingPerDay := strings.ToLower(elem.WorkingPerDay)
		salary :=elem.Salary
		skills :=elem.Skills
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






if len(searchWorkingPerDay) > 0 {
convertedSearchWorkingPerDay, err :=  strconv.Atoi(searchWorkingPerDay)
fmt.Println("PARTSSS")
parts := strings.Split(workingPerDay, "-")
fmt.Println(parts)
	secParts := strings.Split(parts[1], " ")
	fmt.Println(secParts)
	convertedWorkingPerDay, err  := strconv.Atoi(secParts[0])
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "error"})
	}
	booleanValue:= isLessTimeThanRequired(convertedWorkingPerDay, convertedSearchWorkingPerDay)
	if booleanValue == true {
	} else {
		isAbleToAddIntoFilterdItems=false
	} 
}


if searchTo > 0 {
	parts := strings.Split(salary, "-")  // промежуток из бд
	
first, err := strconv.Atoi(parts[0])
second, err :=strconv.Atoi(parts[1])
if err != nil {
	fmt.Println("ERRRRRRRRRRRRR")
	fmt.Println(err)
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "error"})
}
	//if containsTitle(location, searchLocation) {

		fmt.Println("from ") 
		fmt.Println(searchFrom, first)
		fmt.Println("to")
		fmt.Println(searchTo, second)
		if first<=searchFrom && second>=searchTo  {

	} else {
		isAbleToAddIntoFilterdItems=false
	}
}




 
if len(skills) > 0 {
	skillsLower := make([]string, len(skills))
	for i, v := range skills {
	skillsLower[i] = strings.ToLower(v)
	}
	



//	allPresent := true
    for _, v := range searchSkillsLower {
        found := false
        for _, item := range skillsLower {
            if item == v {
                found = true
                break
            }
        }
        if !found {
			isAbleToAddIntoFilterdItems=false
           // allPresent = false
            break
        }
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
	return c.JSON(http.StatusOK, filteredOffers) */
	return c.JSON(http.StatusOK, "test")
}
func isLessTimeThanRequired(title, searchTitle int) bool{
return title>=searchTitle
}
func containsTitle(title, searchTitle string) bool {
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

	/*
	    title: PagProps.title,
    skills: PagProps.skills,
    workingPerDay: PagProps.workingPerDay,
    location: PagProps.location,
    from:  PagProps.salary.from,
    to: PagProps.salary.to
	*/
/*package controller

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	e "server/middleware"
	m "server/db"
)

 

func GetYourOffers(c echo.Context) error {
	token := c.QueryParam("token")
	fmt.Println("Received token:", token)
decodedToken, errToken := e.Decode(token,  "key")
fmt.Println("Decoded")
fmt.Println(decodedToken)
fmt.Println("USERNAME"+decodedToken.Username)
fmt.Println(errToken)

user, userErr := m.FindUserByUsername(decodedToken.Username)
	fmt.Println(user)

	if userErr != nil {
		errorMessage := fmt.Sprintf("Error finding user: %s", userErr.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{"Error": errorMessage,
})
	}
fmt.Println("HELOOOOOOOOOOOOOOOOOOO")
fmt.Println(user.YourOffers)
fmt.Println(user.YourOffers[0])
	return c.JSON(http.StatusOK, "h")
} */


/*
package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	e "server/middleware"
	m "server/db"
)
type VacancyData struct {
	ID                int      `json:"id"`
	Title             string   `json:"title"`
	Description       string   `json:"description"`
	Skills            []string `json:"skills"`
	WorkingPerDay     string   `json:"workingPerDay"`
	Location          string   `json:"location"`
	Salary            string   `json:"salary"`
	Owner             string   `json:"owner"`
	ImageSet          []string `json:"image_set"`
	DataOfPublication string   `json:"data_of_publication"`
	LastTimeOfRise    string   `json:"last_time_of_rise"`
}
func GetYourOffers(c echo.Context) error {
	token := c.QueryParam("token")
	fmt.Println("Received token:", token)
	decodedToken, errToken := e.Decode(token, "key")
	fmt.Println("Decoded")
	fmt.Println(decodedToken)
	fmt.Println("USERNAME" + decodedToken.Username)
	fmt.Println(errToken)
	user, userErr := m.FindUserByUsername(decodedToken.Username)
	fmt.Println(user)
	if userErr != nil {
		errorMessage := fmt.Sprintf("Error finding user: %s", userErr.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{"Error": errorMessage})
	}
	fmt.Println("HELOOOOOOOOOOOOOOOOOOO")
	fmt.Println(user.YourOffers)
	cleanedOffers := make([]string, len(user.YourOffers))
	for i, offer := range user.YourOffers {
		cleanedOffer := strings.Map(func(r rune) rune {
			if r == '{' || r == '}' {
				return -1  
			}
			return r
		}, offer)
		cleanedOffers[i] = cleanedOffer
	}
	var vacancyDataArray []VacancyData
	for i, offer := range cleanedOffers {
//item, err := m.GetVacancyData(string(offer))
fmt.Println(i)
fmt.Println("OFFER") 
fmt.Println(offer)
item, err := m.FindOfferById(string(offer))
vacancyDataArray = append(vacancyDataArray, item)
fmt.Println("ITEM")
fmt.Println(item)
fmt.Println(err)
}
	return c.JSON(http.StatusOK, cleanedOffers)
}   */


package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	e "server/middleware"
	m "server/db"
)

type VacancyData struct {
	ID                int      `json:"id"`
	Title             string   `json:"title"`
	Description       string   `json:"description"`
	Skills            []string `json:"skills"`
	WorkingPerDay     string   `json:"workingPerDay"`
	Location          string   `json:"location"`
	Salary            string   `json:"salary"`
	Owner             string   `json:"owner"`
	ImageSet          []string `json:"image_set"`
	DataOfPublication string   `json:"data_of_publication"`
	LastTimeOfRise    string   `json:"last_time_of_rise"`
}

func GetYourOffers(c echo.Context) error {
	token := c.QueryParam("token")
	fmt.Println("Received token:", token)
	decodedToken, errToken := e.Decode(token, "key")
	fmt.Println("Decoded")
	fmt.Println(decodedToken)
	fmt.Println("USERNAME" + decodedToken.Username)
	fmt.Println(errToken)
	user, userErr := m.FindUserByUsername(decodedToken.Username)
	fmt.Println(user)
	if userErr != nil {
		errorMessage := fmt.Sprintf("Error finding user: %s", userErr.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{"Error": errorMessage})
	}
	fmt.Println("HELOOOOOOOOOOOOOOOOOOO")
	fmt.Println(user.YourOffers)
	cleanedOffers := make([]string, len(user.YourOffers))
	for i, offer := range user.YourOffers {
		cleanedOffer := strings.Map(func(r rune) rune {
			if r == '{' || r == '}' {
				return -1
			}
			return r
		}, offer)
		cleanedOffers[i] = cleanedOffer
	}
	var vacancyDataArray []VacancyData
	for _, offer := range cleanedOffers {
		item, err := m.FindOfferById(offer)
		if err != nil {
			fmt.Printf("Error retrieving offer with ID %s: %s\n", offer, err.Error())
			continue
		}
		vacancyData := VacancyData{
			ID:                item.ID,
			Title:             item.Title,
			Description:       item.Description,
			Skills:            item.Skills,
			WorkingPerDay:     item.WorkingPerDay,
			Location:          item.Location,
			Salary:            item.Salary,
			Owner:             item.Owner,
			ImageSet:          item.ImageSet,
			DataOfPublication: item.DataOfPublication,
			LastTimeOfRise:    item.LastTimeOfRise,
		}
		vacancyDataArray = append(vacancyDataArray, vacancyData)
	}

	return c.JSON(http.StatusOK, vacancyDataArray)
}

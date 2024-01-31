package controller

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	e "server/middleware"
 	"database/sql"
	 m "server/db"

	//m "server/db"
)
 var (
	DB *sql.DB
)
 type UserData struct {
	Username             string `json:"username"`
	Country              string `json:"country"`
	City                 string `json:"city"`
	Telephone            string `json:"telephone"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	Role                 string `json:"role"`
	RegistrationData     string `json:"registration_data"`
	Avatar               string `json:"avatar"`
	Document             string `json:"document"`
	FavouriteOffers      string `json:"favourite_offers"`
	Experience           string `json:"experience"`
	Education            string `json:"education"`
	LastTimeAtNetwork    string `json:"last_time_at_network"`
	Chats                string `json:"chats"`
	Describtion   string `json:"describtion"`
   User_id               string `json:"user_id"`
}
func GetPersonalInformation(c echo.Context) error {
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
	resObj := map[string]interface{}{
        "username": user.Username,
		"country": user.Country, 
		"city": user.City,
		"telephone": user.Telephone,
		"email": user.Email,
		"RegistrationData": user.RegistrationData,
		"avatar": user.Avatar,
		"document": user.Document,
		"favouriteOffers":  user.FavouriteOffers,
		"experience": user.Experience,
		"lastTimeAtNetwork": user.LastTimeAtNetwork,
		"education": user.Education,
		"describtion": user.Describtion,
		 

    }
	fmt.Println("User found:", user)  
	return c.JSON(http.StatusOK,resObj)
}
 
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
/*
func FindUserByUsername(username string) (UserData, error) {
	fmt.Println("EMAILLLLLLLLLLLLLLLLLL")
	fmt.Println(username)
	if DB == nil {
		return UserData{}, fmt.Errorf("Database connection is not established. Call Connect function first.")
	}

	query := "SELECT * FROM user_data WHERE username = $1"
	row := DB.QueryRow(query, username)

	var user UserData
	err := row.Scan(
		&user.Username,
		&user.Country,
		&user.City,
		&user.Telephone,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.RegistrationData,
		&user.Avatar,
		&user.Document,
		&user.FavouriteOffers,
		&user.Experience,
		&user.Education,
		&user.LastTimeAtNetwork,
		&user.Chats,
		&user.Describtion,
		&user.User_id,
	)

	if err == sql.ErrNoRows {
		fmt.Println(user)
		return UserData{}, fmt.Errorf("User not found with email: %s", username)
	} else if err != nil {
		return UserData{}, fmt.Errorf("Failed to query user_data: %v", err)
	}
fmt.Println(user)
	return user, nil
} */
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

	fmt.Println("User found:", user)  
	return c.JSON(http.StatusOK,user)
}
 
package controller

import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/labstack/echo/v4"
)
type RegistrationParams struct {
	Username  string `json:"username"`
	Country   string `json:"country"`
	City      string `json:"city"`
	Telephone string `json:"telephone"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func Register(c echo.Context) error {
	//u:="Helooo"
	fmt.Println("Is workink")
	 


	var registrationData RegistrationParams
	err := json.NewDecoder(c.Request().Body).Decode(&registrationData)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// Далее вы можете использовать registrationData в вашем коде

	fmt.Println("Received registration data:", registrationData)

	return c.JSON(http.StatusOK, map[string]string{"message": "Registration successful"})
 // return c.JSON(http.StatusOK, u)



}

package controller
import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/labstack/echo/v4"
	 "server/db"

	 m "server/middleware"
)
type RegistrationParams struct {
	Username  string `json:"username"`
	Country   string `json:"country"`
	City      string `json:"city"`
	Telephone string `json:"telephone"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Code  string `json:"code"`
}
func Register(c echo.Context) error {
	fmt.Println("CODE IN REGISTER"+code)
    fmt.Println("Is working")
    var registrationData RegistrationParams
    err := json.NewDecoder(c.Request().Body).Decode(&registrationData)
    if err != nil {
        fmt.Println("Error decoding request body:", err)
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
    }
    fmt.Println("Received registration data:", registrationData)
	hashed, hashErr:=m.Encode(registrationData.Password, 180, "key")
	fmt.Println(hashed)
	fmt.Println(hashErr)
	fmt.Println("CODEEEEEEEEEEEEEEEEEEEEEE"+registrationData.Code)
	fmt.Println("CODE "+code)
	fmt.Println("Reg CODE "+registrationData.Code)
	if(registrationData.Code==code){
		fmt.Println("Reg paramssssssssssssssssssssssssssssssss")
		fmt.Println(registrationData.Username)
		fmt.Println(registrationData.City)
		fmt.Println(registrationData.Country)
		fmt.Println(registrationData.Password)
		fmt.Println(registrationData.Username)
		fmt.Println(registrationData.Telephone)
		fmt.Println(registrationData.Email)
		err = db.InsertData(c.Response(), registrationData.Username, registrationData.Password, registrationData.Country, registrationData.City, registrationData.Telephone, registrationData.Email)
		return c.JSON(http.StatusOK, map[string]string{"finalRegisterMessage": "correct",})
	}
    if err != nil {
        fmt.Println("Error at registration:", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Registration failed"})
    }
	return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Registration failed"})
}


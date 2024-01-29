/*package controller
import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/labstack/echo/v4"
	m "server/db"
	e "server/middleware"
)
type LoginParams struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
}
func Login(c echo.Context) error {
	var loginData LoginParams
    err := json.NewDecoder(c.Request().Body).Decode(&loginData)
    if err != nil {
        fmt.Println("Error decoding request body:", err)
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
    }
    fmt.Println("Received login data:", loginData)
user, userErr:=m.FindUserByEmail(loginData.Email)
fmt.Println(user)
if(userErr!=nil){
	return c.JSON(http.StatusBadRequest, "something went wrong")
}
fmt.Println("err")
jwt, jwtErr:=e.Encode(user.User_id, 180, "key")
if(jwtErr!=nil) {
   fmt.Println((jwtErr))
}
return c.JSON(http.StatusOK, jwt)
}  */


package controller

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/labstack/echo/v4"
 
	m "server/db"
	e "server/middleware"
)

type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c echo.Context) error {
	var loginData LoginParams
	err := json.NewDecoder(c.Request().Body).Decode(&loginData)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"Error": "Error generating JWT"})
	}

	fmt.Println("Received login data:", loginData)

	user, userErr := m.FindUserByEmail(loginData.Email)
	fmt.Println(user)

	if userErr != nil {
		errorMessage := fmt.Sprintf("Error finding user: %s", userErr.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{"Error": errorMessage})
	}

	fmt.Println("User found:", user)

	jwt, jwtErr := e.Encode(user.User_id, 180, "key")
	if jwtErr != nil {
		fmt.Println(jwtErr)
		return c.JSON(http.StatusInternalServerError, map[string]string{"Error": "Error generating JWT"})
	}
fmt.Println("password" +loginData.Password)
fmt.Println("email "+user.Email)
//hasheassword := hashedPassword(loginData.Password)
if(user.Password==loginData.Password){

	return c.JSON(http.StatusOK, jwt)
} else {
	return c.JSON(http.StatusInternalServerError, map[string]string{"Error": "Error generating JWT"})
}
}

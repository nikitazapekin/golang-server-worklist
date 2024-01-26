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
if(userErr!=nil){
	fmt.Println("err")
}
 jwt, jwtErr:=e.Encode(user.User_id, 180, "key")
 if(jwtErr!=nil) {
	fmt.Println((jwtErr))
 }
 return c.JSON(http.StatusOK, jwt)
}
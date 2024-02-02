
package controller
import (
	"fmt"
	"strings"
	"regexp"
	"math/rand"
	//gomail "gopkg.in/gomail.v2"
	"net/http"
	"encoding/json"
	"github.com/labstack/echo/v4"
	m "server/db"
	"time"
)
func generateNumber() string {
	rand.Seed(time.Now().UnixNano())

	var result string
	for i := 0; i < 6; i++ {
		digit := rand.Intn(10)
		result += fmt.Sprint(digit)
	}

	return result
}
var code string = "ewww"
func isValidPassword(password string) bool {
	return len(password) >= 6 && strings.ContainsAny(password, "0123456789")
}

func isValidEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(regex).MatchString(email)
}

func SendCodeToEmail(c echo.Context) error {
	fmt.Println(code)
	fmt.Println("worl")
	var regData RegistrationParams
	err := json.NewDecoder(c.Request().Body).Decode(&regData)
	fmt.Println("EMAIL" + regData.Email)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}
	if !isValidEmail(regData.Email) {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid email address"})
	}
	if !isValidPassword(regData.Password) {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid password. It should be at least 6 characters and contain at least 1 digit"})
	}
	if user, err := m.FindUserByEmail(regData.Email); err == nil {
		fmt.Println(user)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User already registered with this email"})
	}
	
	code=generateNumber()
	fmt.Println("CODEEEEEEEEEEE THAT UOU NEED TO ENTER" +code)
	return c.JSON(http.StatusOK, "Please type code that was sended on "+ regData.Email +": "+code)
}


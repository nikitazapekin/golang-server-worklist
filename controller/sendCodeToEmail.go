
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
	/*key := "xsmtpsib-e6ed306964a46e88f2c96db9d3d09bb4406a08f87ba9066c2ec61665dd206d6d-IO3F9rwdmjSgARp1"
	from := "testemailforprojects341@gmail.com"
	host := "smtp-relay.brevo.com"
	port := 587
	to := "dvijhfxr@gmail.com"
	msg := gomail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", "HElo")
	msg.SetBody("text/plain", "helo")
	n := gomail.NewDialer(host, port, from, key)
	if err := n.DialAndSend(msg); err != nil {

		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to send email"})
	} */
	code=generateNumber()
	
	return c.JSON(http.StatusOK, "Please type code that was sended on "+ regData.Email +": "+code)
}


/*
package controller

import (
	"fmt"
	gomail "gopkg.in/gomail.v2"
	"net/http"
	"encoding/json"
	"github.com/labstack/echo/v4"
)


var code string="ewww"
func SendCodeToEmail(c echo.Context) error {
	fmt.Println(code)
	fmt.Println("worl")
 
	u := &User{
		Name:  "Jon",
		Email: "jon@labstack.com",
	}

	var regData RegistrationParams
	err := json.NewDecoder(c.Request().Body).Decode(&regData)
	fmt.Println("EMAIL"+regData.Email)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}




key :="xsmtpsib-e6ed306964a46e88f2c96db9d3d09bb4406a08f87ba9066c2ec61665dd206d6d-IO3F9rwdmjSgARp1" 
from :="testemailforprojects341@gmail.com"
host:="smtp-relay.brevo.com"
port:=587
to:="dvijhfxr@gmail.com"
msg := gomail.NewMessage()
msg.SetHeader("From", from)
msg.SetHeader("To", to)
msg.SetHeader("Subject", "HElo")
msg.SetBody("text/plain", "helo")
n:=gomail.NewDialer(host, port, from, key)
if err := n.DialAndSend(msg); err!=nil {
	//panic(err)
}
	return c.JSON(http.StatusOK, u)
} 
 */
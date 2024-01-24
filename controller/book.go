package controller

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetBooks(c echo.Context) error {
	u := &User{
		Name:  "Jon",
		Email: "jon@labstack.com",
	}
fmt.Println("HELOOOOOOOOOOOOOOOOOOO")
	return c.JSON(http.StatusOK, u)
}
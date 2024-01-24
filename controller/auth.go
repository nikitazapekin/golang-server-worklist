package controller

import (
	//"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
)
/*
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
} */

func Register(c echo.Context) error {
	u:="Helooo"
	/*u := &User{
		Name:  "Jon",
		Email: "jon@labstack.com",
	}
  */

  return c.JSON(http.StatusOK, u)



}

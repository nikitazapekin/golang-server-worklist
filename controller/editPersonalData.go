
	package controller

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	"encoding/json"
	e "server/middleware"
	m "server/db"
)

type EditDataParams struct {
	Education   string `json:"education"`
	About   string `json:"about"`
	Experience   string `json:"experience"`
	Email   string `json:"email"`
	Password   string `json:"password"`
	Tepephone   string `json:"telephone"`
	Country   string `json:"country"`
	City   string `json:"city"`
	Document   string `json:"document"`
	Token  string `json:"token"`
}

func EditPersonalInformation(c echo.Context) error {
	var editDataParams EditDataParams
	err := json.NewDecoder(c.Request().Body).Decode(&editDataParams)
	fmt.Println(err)
	fmt.Println(editDataParams)
	fmt.Println((editDataParams.Country))
isCorrectInputData:=true;
	decodedToken, errToken := e.Decode(editDataParams.Token,  "key")
	fmt.Println("Decoded")
	fmt.Println(decodedToken)
	fmt.Println("USERNAME"+decodedToken.Username)
	fmt.Println(errToken)
	fmt.Println(isCorrectInputData)
	fmt.Println("USER"+decodedToken.Username)
	currentUserEmail :=""
	if !isValidEmail(editDataParams.Email) {
		isCorrectInputData=false;
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid email address"})
	}
	if !isValidPassword(editDataParams.Password) {
		isCorrectInputData=false;
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid password. It should be at least 6 characters and contain at least 1 digit"})
	}
	if user, err := m.FindUserByUsername(decodedToken.Username); err == nil {
		fmt.Println(user)
		currentUserEmail=user.Email
		
if(user.Email==editDataParams.Email){
	isCorrectInputData=false;
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Please type another email adress"})
}
userr, errr := m.FindUserByEmail(editDataParams.Email);
if(userr.Email!= "") {
	isCorrectInputData=false;
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "User already registered with this email"}) 
}
fmt.Println(userr, errr)
}
if(isCorrectInputData){
	user, err := m.FindUserByEmail(editDataParams.Email);
	fmt.Println(user, err)
	user.Education=editDataParams.Education  
	user.Describtion =editDataParams.About  
	user.Experience  =editDataParams.Experience  
	user.Email  =editDataParams.Email  
	user.Password =editDataParams.Password  
	user.Telephone =editDataParams.Tepephone 
	user.Country =editDataParams.Country   
	user.City   =editDataParams.City
	user.Document  =editDataParams.Document
//	 m.UpdateUser(user)
errr := m.UpdateUser(user, currentUserEmail)
fmt.Println("ERROR ", errr)
	fmt.Println("Everything is clear")
	fmt.Println(user)
	return c.JSON(http.StatusOK,  user,)
}
return c.JSON(http.StatusBadRequest, map[string]string{"error": "something went wrong"})
}
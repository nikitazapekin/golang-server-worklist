package middleware

import (
	"fmt"
	//"net/http"
//	"os"
//	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
//	"github.com/labstack/echo/v4"
)

var APPLICATION_NAME = "golang-mvc-rest-api"
var LOGIN_EXPIRATION_DURATION int64
var REFRESH_TOKEN_DURATION int64

type JWTClaim struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

func Encode(username string, duration int64, secretKey string) (string, error) {
	fmt.Println("======secret key=======")
	fmt.Printf("%+v\n", secretKey)
	fmt.Printf("%+v\n", duration)
	fmt.Println("=============")
	claims := JWTClaim{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: duration,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

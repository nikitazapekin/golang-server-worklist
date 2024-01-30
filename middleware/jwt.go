package middleware

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var APPLICATION_NAME = "golang-mvc-rest-api"
var LOGIN_EXPIRATION_DURATION int64
var REFRESH_TOKEN_DURATION int64

type JWTClaim struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

/*
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
*/


func Encode(username string, duration int64, secretKey string) (string, error) {
	fmt.Println("======secret key=======")
	fmt.Printf("%+v\n", secretKey)
	fmt.Printf("%+v\n", duration)
	fmt.Println("=============")

	// Используйте текущее время и добавьте к нему duration, чтобы установить срок действия токена.
	expirationTime := time.Now().Add(time.Duration(duration) * time.Second)

	claims := JWTClaim{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: expirationTime.Unix(),
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


func Decode(tokenString string, secretKey string) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaim); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("Invalid token")
}

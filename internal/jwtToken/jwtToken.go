package jwtToken

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var signKey = []byte("95556f1139732990b9f79ac1c4b6572f40195d4e4dcf1e02e7a6d8e2c2931572")

func Create(login string) (string, error){
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = login
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString(signKey)
	if err != nil{
		return "", err
	}
	return tokenString, nil
}
package jwtToken

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func ExtractLogin(tokenString string) (string, error) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure that the token's algorithm is valid (optional but recommended)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if username, ok := claims["username"].(string); ok {
			return username, nil
		}
		return "", fmt.Errorf("username claim not found")
	}

	return "", fmt.Errorf("invalid token")
}

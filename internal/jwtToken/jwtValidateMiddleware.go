package jwtToken

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func CheckLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			http.Error(w, "token is missing", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return signKey, nil
		})
		if err != nil {
			log.Println("error parsing token: ", token, "error: ", err)
		}

		if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
			http.Error(w, "unauthorized token", http.StatusUnauthorized)
		}
		next.ServeHTTP(w, r)
	})
}

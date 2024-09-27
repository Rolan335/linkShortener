package util

import (
	"LinkShortener/internal/db/postgres"
	"LinkShortener/internal/jwtToken"
	"log"
)

func GetClientIdByToken(token string) uint {
	login, err := jwtToken.ExtractLogin(token)
	if err != nil {
		log.Println("error extracting login from jwt: ", login, "error: ", err)
	}
	user := postgres.Clients{}
	postgres.Db.Select("id").Find(&user, "login = ?", login)
	return user.ID
}

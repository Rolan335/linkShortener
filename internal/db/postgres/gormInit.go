package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect(host string, user string, password string, dbName string, port uint) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", host, user, password, dbName, port)
	fmt.Println(password)
	fmt.Printf("host=%s user=%s password=%s dbname=%s port=%d", host, user, password, dbName, port)
	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	Db.AutoMigrate(&Clients{}, &Links{})
	return nil
}

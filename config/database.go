package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadDB() {

	connectionStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC", ENV.DB_HOST, ENV.DB_PORT, ENV.DB_USER, ENV.DB_PASSWORD, ENV.DB_NAME)
	db, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db
}
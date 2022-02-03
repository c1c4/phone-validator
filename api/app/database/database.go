package database

import (
	"api/app/config"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() {
	var err error
	Database, err = gorm.Open(sqlite.Open(config.DATABASE), &gorm.Config{})

	if err != nil {
		log.Fatal("it's not possible to connect with the database", err)
	}
}

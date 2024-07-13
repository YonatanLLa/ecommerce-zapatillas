package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost user=postgres password=admin123 dbname=api_golang port=5432"
var DB *gorm.DB

func DBConnect() {
	var err error
	// Connect to database
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Database connected")
	}
}

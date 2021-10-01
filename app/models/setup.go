package models

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	hostname := os.Getenv("DATABASE_HOST")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	databaseName := os.Getenv("DATABASE_NAME")
	port := os.Getenv("DATABASE_PORT")

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", hostname, user, databaseName, port, password)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		QueryFields: true,
	})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(dbModels...)
	if err != nil {
		log.Fatal(err)
	}

	DB = database
}

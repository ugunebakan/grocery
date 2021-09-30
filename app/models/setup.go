package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func ConnectDataBase() {
  dsn := "host=localhost user=admumugun dbname=grocery2 port=5432"
  database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
	QueryFields: true,
  })

  //database, err := gorm.Open("sqlite3", "test.db")

  if err != nil {
    panic("Failed to connect to database!")
  }

  database.AutoMigrate(dbModels...)

  DB = database
}

package models

import (
	"log"
	"os"
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
  "github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
        database, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})

        if err != nil {
                panic("Failed to connect to database!")
        }

        err = database.AutoMigrate(&User{})
        if err != nil {
                log.Fatal("Error migrating DB", err)
        }

        DB = database
}
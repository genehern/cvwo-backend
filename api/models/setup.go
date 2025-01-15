package models

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

        err = database.Debug().AutoMigrate(&User{}, &Post{}, &Comment{})
        if err != nil {
                log.Fatal("Error migrating DB", err)
        }

        DB = database
}
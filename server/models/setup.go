package models

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectionDataBase() {
	// Checking .env file
	godotenv.Load()
	dsn, exist := os.LookupEnv("POSTGRES_CONNECT")

	if !exist {
		log.Printf("Error loading .env file")
	}

	// Database connection
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// Adding schema to database
	database.AutoMigrate()

	DB = database
}

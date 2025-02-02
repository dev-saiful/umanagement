package config

import (
	"fmt"
	"log"
	"os"

	"github.com/dev-saiful/umanagement/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// database connection variable
var DB *gorm.DB

func InitDB() {
	host := os.Getenv("DATABASE_HOST")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	name := os.Getenv("DATABASE_NAME")
	// Postgresql connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", host, user, password, name)
    log.Println("Connecting to database with DSN")
	// Open connection to database
	db, err := gorm.Open((postgres.Open(dsn)), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}
	// Auto migrate models
	db.AutoMigrate(&models.User{})
	DB = db
	fmt.Println("Connection to database established!")
}

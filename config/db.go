package config

import (
	"fmt"

	"github.com/dev-saiful/umanagement/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// database connection variable
var DB *gorm.DB

func InitDB() {
	// Postgresql connection string
	dsn := "host=ep-shy-shape-a1aajhjj.ap-southeast-1.pg.koyeb.app user=koyeb-adm password=npg_Ixl2rvZ9NCKW dbname=koyebdb"

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

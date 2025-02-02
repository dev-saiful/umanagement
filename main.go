package main

import (
	"log"

	"github.com/dev-saiful/umanagement/config"
	"github.com/dev-saiful/umanagement/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }
	config.InitDB()
	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":5000") // listen and serve on 0.0.0.0:5000
}

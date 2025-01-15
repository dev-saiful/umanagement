package main

import (
	"usermanagement/config"
	"usermanagement/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	
	config.InitDB()

	router := gin.Default()
	routes.AuthRoutes(router)

	router.Run(":5000") // listen and serve on 0.0.0.0:5000
}

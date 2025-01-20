package main

import (
	"github.com/dev-saiful/umanagement/config"
	"github.com/dev-saiful/umanagement/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	
	config.InitDB()
	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":5000") // listen and serve on 0.0.0.0:5000
}

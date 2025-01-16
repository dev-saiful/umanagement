package routes

import (
	"usermanagement/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {

	router.POST("/auth/login",controllers.Login)
	router.POST("/auth/signup",controllers.Signup)
	
}
package routes

import (
	"github.com/dev-saiful/umanagement/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	authorized := router.Group("api/v1/auth")
	{
		authorized.POST("/login",controllers.Login)
		authorized.POST("/signup",controllers.Signup)
	}
	
}
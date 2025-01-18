package routes

import (
	"github.com/dev-saiful/umanagement/controllers"
	"github.com/dev-saiful/umanagement/middlewares"
	"github.com/gin-gonic/gin"
)



func UserRoutes(router *gin.Engine) {
	private := router.Group("api/v1/user")
	private.Use(middlewares.AuthMiddleware())
	{
		private.GET("/profile", controllers.GetProfile)	
	}
}
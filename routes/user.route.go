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
		private.GET("/admin",middlewares.RoleMiddleware("admin"),controllers.GetAdmin)
		private.GET("/users",middlewares.RoleMiddleware("admin"),controllers.GetAllUser)	
		private.GET("/users/:id",middlewares.RoleMiddleware("admin"),controllers.GetUserById)
		private.PUT("/users/:id",middlewares.RoleMiddleware("admin"),controllers.UpdateUser)
		private.DELETE("/users/:id",middlewares.RoleMiddleware("admin"),controllers.DeleteUser)

	}
}
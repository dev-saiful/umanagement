package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {

		AuthRoutes(router)
		UserRoutes(router)


}

package controllers

import (
	"net/http"

	"github.com/dev-saiful/umanagement/services"
	"github.com/gin-gonic/gin"
)

func GetProfile(ctx *gin.Context) {
	// Get the user email from the context
	email, _ := ctx.Get("email")
	// Initialize user model
	user, err := services.GetProfile(email.(string))
	// Check if user is not found
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

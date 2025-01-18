package controllers

import (
	"fmt"
	"net/http"

	"github.com/dev-saiful/umanagement/config"
	"github.com/dev-saiful/umanagement/models"
	"github.com/gin-gonic/gin"
)

func GetProfile(ctx *gin.Context) {

	// Get the user email from the context
	email, _ := ctx.Get("email")
	fmt.Print(email)
	// Initialize user model
	var user models.User
	db := config.DB
	// Query the database for the user by email
	err := db.Where("email = ?", email).First(&user).Error

	// Check if user is not found
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

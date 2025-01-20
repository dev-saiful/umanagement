package controllers

import (
	"net/http"

	"github.com/dev-saiful/umanagement/config"
	"github.com/dev-saiful/umanagement/models"
	"github.com/gin-gonic/gin"
)

func GetAdmin(ctx *gin.Context) {
	// Get the user email from the context
	email, _ := ctx.Get("email")
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

func GetAllUser(ctx *gin.Context) {
	var users []models.User
	db := config.DB
	err := db.Select("id,email,username").Find(&users).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve users"})
		return
	}
	if len(users) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No users found"})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func GetUserById(ctx *gin.Context) {
	// Get the user email from the context
	userId := ctx.Param("id")
	// Initialize user model
	var user models.User
	db := config.DB
	// Query the database for the user by email
	err := db.First(&user, userId).Error

	// Check if user is not found
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func UpdateUser(ctx *gin.Context) {
	// Get the user ID from the URL parameter
	userId := ctx.Param("id")

	// Initialize user model
	var user models.User
	db := config.DB

	// Query the database for the user by ID
	err := db.First(&user, userId).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	// Bind the JSON payload to the user model
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	// Save the updated user record to the database
	err = db.Save(&user).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUser(ctx *gin.Context) {
	// Get the user ID from the URL parameter
	userId := ctx.Param("id")

	// Initialize user model
	var user models.User
	db := config.DB

	// Query the database for the user by ID
	err := db.First(&user, userId).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	// Delete the user record from the database
	err = db.Delete(&user).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

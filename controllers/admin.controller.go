package controllers

import (
	"net/http"
	"strconv"

	"github.com/dev-saiful/umanagement/models"
	"github.com/dev-saiful/umanagement/services"
	"github.com/gin-gonic/gin"
)

func GetAdmin(ctx *gin.Context) {
	// Get the user email from the context
	email, _ := ctx.Get("email")
	user, err := services.GetAdmin(email.(string))

	// Check if user is not found
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func GetAllUser(ctx *gin.Context) {
	users, err := services.GetAlluser()
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
	// Get the user ID from the URL parameter
	userId, _ := strconv.Atoi(ctx.Param("id"))
	user, err := services.GetUserById(userId)
	// Check if user is not found
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func UpdateUser(ctx *gin.Context) {
	// Get the user ID from the URL parameter
	userId, _ := strconv.Atoi(ctx.Param("id"))
	var user models.User
	// Bind the JSON payload to the user model
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}
	err := services.UpdateUser(userId, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUser(ctx *gin.Context) {
	// Get the user ID from the URL parameter
	userId, _ := strconv.Atoi(ctx.Param("id"))
	err := services.DeleteUser(userId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

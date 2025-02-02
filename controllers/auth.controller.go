package controllers

import (
	"net/http"

	"github.com/dev-saiful/umanagement/models"
	"github.com/dev-saiful/umanagement/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func Login(ctx *gin.Context) {
	var loginReq models.LoginRequest
	// Bind the incoming JSON request to loginReq struct
	err := ctx.ShouldBindJSON(&loginReq)
	if err != nil {
		// If binding fails, return a 400 Bad Request response with an error message
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "All fields are required"})
		return
	}
	// Validate the Email field to ensure it is in a valid email format
	err = validate.Var(loginReq.Email, "required,email")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email"})
		return
	}

	token, err := services.Login(loginReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Return a 200 OK response indicating the user was successfully logged in
	ctx.JSON(http.StatusOK, gin.H{"message": "Login", "token": token})
}

func Signup(ctx *gin.Context) {
	var signupReq models.SignupRequest
	// Bind the incoming JSON request to signupReq struct
	err := ctx.ShouldBindJSON(&signupReq)
	if err != nil {
		// If binding fails, return a 400 Bad Request response with an error message
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "All fields are required"})
		return
	}

	// Validate the Username field to ensure it's between 4 and 32 characters
	err = validate.Var(signupReq.Username, "required,min=4,max=32")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Username must be between 4 and 32 characters"})
		return
	}

	// Validate the Email field to ensure it is in a valid email format
	err = validate.Var(signupReq.Email, "required,email")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email"})
		return
	}

	// Validate the Password field to ensure it is at least 6 characters long
	err = validate.Var(signupReq.Password, "required,min=6")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Password must be at least 6 characters"})
		return
	}

	// Ensure the Password and ConfirmPassword fields match
	err = validate.VarWithValue(signupReq.Password, signupReq.ConfirmPassword, "eqfield")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Passwords do not match"})
		return
	}

	// Save the new user record to the database
	err = services.Signup(signupReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Return a 201 Created response indicating the user was successfully created
	ctx.JSON(http.StatusCreated, gin.H{"message": "Signup successful"})
}

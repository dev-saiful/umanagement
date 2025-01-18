package controllers

import (
	"net/http"

	"github.com/dev-saiful/umanagement/config"
	"github.com/dev-saiful/umanagement/models"
	"github.com/dev-saiful/umanagement/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func Login(ctx *gin.Context) {
	var db = config.DB 
	var loginReq models.LoginRequest 
	err := ctx.ShouldBindJSON(&loginReq) // Bind the incoming JSON request to loginReq struct
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
	// user exists
	var user models.User
	err = db.Where("email = ?", loginReq.Email).First(&user).Error
if err != nil{
	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
        return
}
	// Check if the provided password matches the hashed password in the database
	err = user.CheckPassword(loginReq.Password)
	if err != nil {
		// If the passwords do not match, return a 401 Unauthorized response
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email or password"})
		return
	}
	// Generate JWT token
	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token"})
		return
	}

	// Return a 200 OK response indicating the user was successfully logged in
	ctx.JSON(200, gin.H{"message": "Login","token": token})
}

func Signup(ctx *gin.Context) {
	var db = config.DB            
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


	// Create a new User instance with the validated data
	user := models.User{
		Username: signupReq.Username,
		Password: signupReq.Password, // The password will be hashed before saving
		Email:    signupReq.Email,
	}
	// Hash the user's password before storing it in the database
	err = user.HashPassword(user.Password)
	if err != nil {
		// If password hashing fails, return a 500 Internal Server Error response
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to hash password"})
		return
	}

	// Save the new user record to the database
	err = db.Create(&user).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Email or username already exists"})
		return
	}

	// Return a 201 Created response indicating the user was successfully created
	ctx.JSON(http.StatusCreated, gin.H{"message": "Signup successful"})
}


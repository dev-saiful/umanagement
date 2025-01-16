package controllers

import (
	"errors"
	"net/http"
	"usermanagement/config"
	"usermanagement/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func Login(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Login"})
}

func Signup(ctx *gin.Context) {
	var db = config.DB
	var signupReq models.SignupRequest
	// request body
	err := ctx.ShouldBindJSON(&signupReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "All fileds are required"})
		return
	}
	// Validate the input
	err = validate.VarWithValue(signupReq.Password, signupReq.ConfirmPassword, "eqfield")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Passwords do not match"})
		return
	}

	// create user
	user := models.User{
		Username: signupReq.Username,
		Password: signupReq.Password,
		Email:    signupReq.Email,
	}
	// checking if user already exists
	err = db.Where("email = ?", user.Email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// proceed
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Database error"})
		return
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "User already exists"})
		return
	}

	// hash the password
	err = user.HashPassword(user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to hash password"})
		return
	}

	// save the user to the database
	err = db.Create(&user).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Signup"})
}

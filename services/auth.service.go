package services

import (
	"errors"

	"github.com/dev-saiful/umanagement/config"
	"github.com/dev-saiful/umanagement/models"
	"github.com/dev-saiful/umanagement/utils"
)

func Login(loginReq models.LoginRequest) (string, error) {
	db := config.DB
	var user models.User
	// Check if the user exists in the database
	if err := db.Where("email = ?", loginReq.Email).First(&user).Error; err != nil {
		return "", errors.New("invalid email or password")
	}

	// Verify the password
	if err := utils.CheckPassword(user.Password, loginReq.Password); err != nil {
		return "", errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.Email, user.Role)
	if err != nil {
		return "", errors.New("could not generate token")
	}

	return token, nil
}

func Signup(signupReq models.SignupRequest) error {
	db := config.DB
	// Hash the user's password
	hashpass, err := utils.HashPassword(signupReq.Password)
	if err != nil {
		return errors.New("failed to hash password")
	}

	// Create a new user instance
	user := models.User{
		Username: signupReq.Username,
		Password: hashpass,
		Email:    signupReq.Email,
		Role:     signupReq.Role,
	}

	// Save the new user in the database
	if err := db.Create(&user).Error; err != nil {
		return errors.New("email or username already exists")
	}

	return nil
}

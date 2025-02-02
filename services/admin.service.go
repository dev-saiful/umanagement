package services

import (
	"errors"

	"github.com/dev-saiful/umanagement/config"
	"github.com/dev-saiful/umanagement/models"
)

func GetAdmin(email string) (*models.User, error) {
	// Initialize user model
	var user models.User
	db := config.DB
	err := db.Select("id,email,username").Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetAlluser() ([]models.User, error) {
	var users []models.User
	db := config.DB
	err := db.Select("id,email,username,role").Find(&users).Error
	return users, err
}

func GetUserById(id int) (*models.User, error) {
	var user models.User
	db := config.DB
	err := db.Select("id,email,username,role").First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUser(id int, updatedUser *models.User) error {
	var user models.User
	db := config.DB
	err := db.Select("id,email,username,role").First(&user, id).Error
	if err != nil {
		return errors.New("user not found")
	}
	user.Username = updatedUser.Username
	user.Email = updatedUser.Email
	user.Role = updatedUser.Role

	return db.Save(&user).Error
}

func DeleteUser(id int) error {
	var user models.User
	db := config.DB
	err := db.First(&user, id).Error
	if err != nil {
		return errors.New("user not found")
	}
	return db.Delete(&user).Error
}

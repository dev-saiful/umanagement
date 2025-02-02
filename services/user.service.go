package services

import (
	"github.com/dev-saiful/umanagement/config"
	"github.com/dev-saiful/umanagement/models"
)

func GetProfile(email string) (*models.User, error) {
	db := config.DB
	var user models.User
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

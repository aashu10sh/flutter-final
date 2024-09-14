package usecases

import (
	"fmt"

	auth_entities "github.com/aashu10sh/authapi/modules/auth/domain/entities/dto"
	"github.com/aashu10sh/authapi/modules/core/database/models"
	"gorm.io/gorm"
)

func GetUser(db *gorm.DB, id uint) (models.UserModel, error) {
	var user models.UserModel
	result := db.First(&user, id)
	if result.Error != nil {
		fmt.Println("Error is Happening")
		return user, result.Error
	}
	return user, nil
}

func CreateUser(db *gorm.DB, user auth_entities.UserWithPassword) error {
	result := db.Create(&models.UserModel{
		Name:     user.Name,
		Password: user.Password,
		Username: user.Username,
		IsAdmin:  false,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetUserWithCondition(db *gorm.DB, condition models.UserModel) (models.UserModel, error) {
	var user models.UserModel
	result := db.Where(&condition).First(&user)

	if result.Error != nil {
		fmt.Println("Error is Happening")
		return user, result.Error
	}
	return user, nil
}

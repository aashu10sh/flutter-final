package auth_controller

import (
	"errors"
	"fmt"

	auth_entities "github.com/aashu10sh/authapi/modules/auth/domain/entities/dto"
	"github.com/aashu10sh/authapi/modules/auth/domain/usecases"
	auth_utils "github.com/aashu10sh/authapi/modules/auth/utils"
	"github.com/aashu10sh/authapi/modules/core/database/models"
	"gorm.io/gorm"
)

type AuthController struct {
	DB gorm.DB
}

func (c *AuthController) Login(user auth_entities.UserWithPassword) (string, error) {
	userDB, err := usecases.GetUserWithCondition(&c.DB, models.UserModel{
		Username: user.Username,
	})
	fmt.Println("User From the Database is")
	fmt.Println(userDB)

	if err != nil {
		if err.Error() == "record not found" {
			return "", errors.New("invalid username or password")
		}
		return "", err

	}
	correct := auth_utils.CheckPasswordHash(user.Password, userDB.Password)
	if !correct {
		return "", errors.New("incorrect credentials")
	}

	return auth_utils.SignJWTToken(user.Username)

}

func (c *AuthController) Register(user auth_entities.UserWithPassword) (string, error) {
	_, err := usecases.GetUserWithCondition(&c.DB, models.UserModel{Username: user.Username})

	if err == nil {
		return "", errors.New("exists")
	}
	hashedPassword, error := auth_utils.HashPassword(user.Password)
	if error != nil {
		return "", errors.New("hash")
	}

	err = usecases.CreateUser(&c.DB, auth_entities.UserWithPassword{
		Name:     user.Name,
		Password: hashedPassword,
		Username: user.Username,
	})
	if err != nil {
		return "", errors.New("failed")
	}
	return auth_utils.SignJWTToken(user.Username)
}

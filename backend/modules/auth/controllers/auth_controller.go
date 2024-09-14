package auth_controller

import (
	"errors"

	auth_entities "github.com/aashu10sh/authapi/modules/auth/domain/entities/dto"
	"github.com/aashu10sh/authapi/modules/auth/domain/usecases"
	auth_utils "github.com/aashu10sh/authapi/modules/auth/utils"
	"github.com/aashu10sh/authapi/modules/core/database/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthController struct {
	DB gorm.DB
}

func (c *AuthController) Login(ctx *fiber.Ctx) {

}

func (c *AuthController) Register(user auth_entities.UserWithPassword) (string, error) {
	_, err := usecases.GetUserWithCondition(&c.DB, models.UserModel{Username: user.Username})

	if err == nil {
		return "", errors.New("user already exists")
	}
	err = usecases.CreateUser(&c.DB, auth_entities.UserWithPassword{
		Name:     user.Name,
		Password: user.Password,
		Username: user.Username,
	})
	if err != nil {
		return "", errors.New("user cannot be created")
	}
	return auth_utils.SignJWTToken(user.Username)
}

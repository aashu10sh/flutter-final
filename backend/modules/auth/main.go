package auth

import (
	"fmt"

	auth_controller "github.com/aashu10sh/authapi/modules/auth/controllers"
	auth_entities "github.com/aashu10sh/authapi/modules/auth/domain/entities/dto"
	auth "github.com/aashu10sh/authapi/modules/auth/domain/entities/request"
	auth_response "github.com/aashu10sh/authapi/modules/auth/domain/entities/response"
	"github.com/aashu10sh/authapi/modules/core/database"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	bind "github.com/idan-fishman/fiber-bind"
)

func AuthRouter(router fiber.Router) {
	auth_controller := auth_controller.AuthController{
		DB: *database.GetDB(),
	}

	validate := validator.New()

	router.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Welcome Traveler, Be Authorized Here:) !")
	})

	router.Post("/register", bind.New(bind.Config{
		Validator: validate,
		Source:    bind.JSON,
	}, &auth.RegisterRequest{}), func(ctx *fiber.Ctx) error {

		registerUser := ctx.Locals(bind.JSON).(*auth.RegisterRequest)

		registered, error := auth_controller.Register(
			auth_entities.UserWithPassword{Name: registerUser.Name, Password: registerUser.Password, Username: registerUser.UserName},
		)
		if error != nil {
			fmt.Println(error)
			switch error.Error() {
			case "exists":
				fmt.Println("hit here!")
				return ctx.Status(fiber.StatusNotFound).JSON(error.Error())
			default:
				return ctx.Status(fiber.StatusBadRequest).JSON(error.Error())
			}
		}
		data := auth_response.TokenResponse{
			Token: registered,
		}
		return ctx.JSON(data)
	})

	router.Post("/login", bind.New(bind.Config{
		Validator: validate,
		Source:    bind.JSON,
	}, &auth.LoginRequest{}), func(ctx *fiber.Ctx) error {
		loginRequest := ctx.Locals(bind.JSON).(*auth.LoginRequest)
		loggedIn, error := auth_controller.Login(
			auth_entities.UserWithPassword{Username: loginRequest.UserName, Password: loginRequest.Password},
		)
		if error != nil {
			fmt.Println(error)
			return ctx.Status(fiber.StatusBadRequest).JSON(error.Error())
		}
		return ctx.JSON(loggedIn)
	})
}

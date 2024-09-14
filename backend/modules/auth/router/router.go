package auth

import "github.com/gofiber/fiber/v2"

func AuthRouter(router fiber.Router) {
	router.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("")
	})
}

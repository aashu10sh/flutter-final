package main

import (
	"log"

	"github.com/aashu10sh/authapi/modules/auth"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	Version1Router := app.Group("/api/v1")

	Version1Router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	Version1Router.Group("/auth").Route("", auth.AuthRouter)
	// app.Group("/auth", auth.AuthRouter)

	log.Fatal(app.Listen(":1337"))
}

package main

import (
	"github.com/gofiber/fiber/v2"
	"khot-queue-api/internal/database"
)

func main() {

	database.Connect()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})

	app.Listen(":80")
}

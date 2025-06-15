package main

import (
	"github.com/gofiber/fiber/v2"
	"khot-queue-api/internal/database"
	"khot-queue-api/internal/handlers"
)

func main() {

	database.Connect() // รับการเชื่อมต่อฐานข้อมูล
	api := fiber.New()

	book := api.Group("/book")

	cook := book.Group("cook-to-order-restaurant")
	cook.Get("/GetMenuItem", handlers.GetMenuItem)

	api.Listen(":80")
}

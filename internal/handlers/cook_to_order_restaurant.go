package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"khot-queue-api/internal/database"
	"khot-queue-api/internal/models"
)

func GetMenuItem(ctx *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(ctx.Query("store_id"))
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "invalid store_id"})
	}

	var items []models.MenuItemModel
	cur, err := database.DB.Collection(string(database.MenuItems)).Find(ctx.Context(), bson.M{"store_id": id})
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	defer cur.Close(ctx.Context())

	if err = cur.All(ctx.Context(), &items); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(items)
}

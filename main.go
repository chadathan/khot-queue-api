package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"khot-queue-api/internal/database"
	"time"
)

func main() {

	db, err := database.Connect() // รับการเชื่อมต่อฐานข้อมูล
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	now := time.Now().UTC()
	userID := primitive.NewObjectID()    // ตัวอย่าง ID ของผู้ใช้
	companyID := primitive.NewObjectID() // ตัวอย่าง ID ของบริษัท

	// รายการเมนู
	menuItems := []database.MenuItem{
		{
			ID:          primitive.NewObjectID(),
			Name:        "ข้าวผัดหมู",
			Description: nil, // ตั้งเป็น nil
			Price:       10.00,
			ImageURL:    nil, // ตั้งเป็น nil
			CreatedDate: now,
			CreatedBy:   userID,
			UpdatedDate: now,
			UpdatedBy:   nil,        // ตั้งเป็น nil
			CompanyID:   &companyID, // เปลี่ยนให้เป็น pointer
		},
	}

	data, err := database.ConvertToBsonItems(menuItems)
	_, err = db.Collection(string(database.MenuItems)).InsertMany(context.Background(), data)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})

	app.Listen(":80")
}

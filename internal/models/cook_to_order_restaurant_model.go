package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MenuItemModel struct {
	ID          primitive.ObjectID  `bson:"_id" json:"id"`                  // รหัสของอาหาร (ObjectID)
	Name        string              `bson:"name" json:"name"`               // ชื่อเมนูอาหาร
	Description *string             `bson:"description" json:"description"` // รายละเอียดของเมนู (อาจจะมี)
	Price       float64             `bson:"price" json:"price"`             // ราคาอาหาร
	ImageURL    *string             `bson:"image_url" json:"image_url"`     // URL ของภาพอาหาร
	StoreID     *primitive.ObjectID `bson:"store_id" json:"store_id"`       // ID ของร้าน (ObjectID)
}

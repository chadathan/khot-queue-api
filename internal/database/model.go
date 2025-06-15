package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CollectionName string

const (
	MenuItems CollectionName = "menu_items"
)

type MenuItem struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`                  // รหัสของอาหาร (ObjectID)
	Name        string             `bson:"name" json:"name"`               // ชื่อเมนูอาหาร
	Description *string            `bson:"description" json:"description"` // รายละเอียดของเมนู (อาจจะมี)
	Price       float64            `bson:"price" json:"price"`             // ราคาอาหาร
	ImageURL    *string            `bson:"image_url" json:"image_url"`     // URL ของภาพอาหาร

	CreatedDate time.Time           `bson:"created_date" json:"created_date"` // วันที่สร้าง
	CreatedBy   primitive.ObjectID  `bson:"created_by" json:"created_by"`     // ID ผู้สร้าง (ObjectID)
	UpdatedDate time.Time           `bson:"updated_date" json:"updated_date"` // วันที่อัปเดต
	UpdatedBy   *primitive.ObjectID `bson:"updated_by" json:"updated_by"`     // ID ผู้อัปเดต (ObjectID)
	CompanyID   *primitive.ObjectID `bson:"company_id" json:"company_id"`     // ID ของบริษัท (ObjectID)
}

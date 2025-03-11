package models

import (
	"github.com/google/uuid"
)

type Product struct {
	BaseModel
	Name        string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text"`
	SKU         string    `gorm:"type:varchar(50);unique;not null"`
	Price       float64   `gorm:"type:decimal(10,2);not null"`
	Cost        float64   `gorm:"type:decimal(10,2);not null"`
	Quantity    int       `gorm:"type:integer;not null;default:0"`
	CategoryID  uuid.UUID `gorm:"type:uuid;not null"`

	Category   *Category      `gorm:"foreignKey:CategoryID"`
	Images     []ProductImage `gorm:"foreignKey:ProductID"`
	OrderItems []OrderItem    `gorm:"foreignKey:ProductID"`
}

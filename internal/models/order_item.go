package models

import (
	"github.com/google/uuid"
)

type OrderItem struct {
	BaseModel
	OrderID    uuid.UUID `gorm:"type:uuid;not null"`
	ProductID  uuid.UUID `gorm:"type:uuid;not null"`
	Quantity   int       `gorm:"type:integer;not null"`
	UnitPrice  float64   `gorm:"type:decimal(10,2);not null"`
	TotalPrice float64   `gorm:"type:decimal(12,2);not null"`

	Order   *Order   `gorm:"foreignKey:OrderID"`
	Product *Product `gorm:"foreignKey:ProductID"`
}

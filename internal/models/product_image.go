package models

import (
	"github.com/google/uuid"
)

type ProductImage struct {
	BaseModel
	ProductID    uuid.UUID `gorm:"type:uuid;not null"`
	ImageURL     string    `gorm:"type:varchar(255);not null"`
	IsMain       bool      `gorm:"type:boolean;default:false"`
	DisplayOrder int       `gorm:"type:integer;default:0"`

	Product *Product `gorm:"foreignKey:ProductID"`
}

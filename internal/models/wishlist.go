package models

import (
	"github.com/google/uuid"
)

type Wishlist struct {
	BaseModel
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	ProductID uuid.UUID `gorm:"type:uuid;not null"`

	User    *User    `gorm:"foreignKey:UserID"`
	Product *Product `gorm:"foreignKey:ProductID"`
}

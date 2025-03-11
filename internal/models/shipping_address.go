package models

import (
	"github.com/google/uuid"
)

type ShippingAddress struct {
	BaseModel
	OrderID       uuid.UUID `gorm:"type:uuid;not null;unique"`
	RecipientName string    `gorm:"type:varchar(100);not null"`
	Phone         string    `gorm:"type:varchar(20);not null"`
	ProvinceID    int       `gorm:"type:integer;not null"`
	DistrictID    int       `gorm:"type:integer;not null"`
	SubDistrictID int       `gorm:"type:integer;not null"`
	Address       string    `gorm:"type:varchar(400);not null"`
	ZipCode       string    `gorm:"type:varchar(10);not null"`

	Order       *Order       `gorm:"foreignKey:OrderID"`
	Province    *Province    `gorm:"foreignKey:ProvinceID"`
	District    *District    `gorm:"foreignKey:DistrictID"`
	SubDistrict *SubDistrict `gorm:"foreignKey:SubDistrictID"`
}

package models

import (
	"github.com/google/uuid"
)

type Address struct {
	BaseModel
	ProvinceID    int     `gorm:"type:integer;not null"`
	DistrictID    int     `gorm:"type:integer;not null"`
	SubDistrictID int     `gorm:"type:integer;not null"`
	Address       *string `gorm:"type:varchar(400)"`

	UserID uuid.UUID `gorm:"type:uuid;not null"`

	Province    *Province
	District    *District
	SubDistrict *SubDistrict
}

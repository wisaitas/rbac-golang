package models

import (
	"github.com/google/uuid"
	"github.com/wisaitas/rbac-golang/pkg"
)

type Address struct {
	pkg.BaseModel
	ProvinceID    uuid.UUID `gorm:"type:uuid;not null"`
	DistrictID    uuid.UUID `gorm:"type:uuid;not null"`
	SubDistrictID uuid.UUID `gorm:"type:uuid;not null"`
	Address       *string   `gorm:"type:varchar(400)"`

	UserID uuid.UUID `gorm:"type:uuid;not null"`

	Province    *Province
	District    *District
	SubDistrict *SubDistrict
}

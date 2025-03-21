package models

import (
	"github.com/google/uuid"
	"github.com/wisaitas/rbac-golang/pkg"
)

type SubDistrict struct {
	pkg.BaseModel
	NameTH     string    `gorm:"type:varchar(100)"`
	NameEN     string    `gorm:"type:varchar(100)"`
	DistrictID uuid.UUID `gorm:"type:uuid"`
	PostalCode string    `gorm:"type:varchar(10)"`

	District *District `gorm:"foreignKey:DistrictID"`
}

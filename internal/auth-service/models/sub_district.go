package models

import (
	"github.com/google/uuid"
	"github.com/wisaitas/rbac-golang/pkg"
)

type SubDistrict struct {
	pkg.BaseModel
	NameTH     string    `gorm:"type:varchar(100)" json:"name_th"`
	NameEN     string    `gorm:"type:varchar(100)" json:"name_en"`
	DistrictID uuid.UUID `gorm:"type:uuid" json:"district_id"`
	ZipCode    int       `gorm:"type:integer" json:"zip_code"`

	District *District `gorm:"foreignKey:DistrictID"`
}

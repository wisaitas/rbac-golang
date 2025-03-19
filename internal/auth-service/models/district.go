package models

import (
	"github.com/google/uuid"
	"github.com/wisaitas/rbac-golang/pkg"
)

type District struct {
	pkg.BaseModel
	NameTH     string    `gorm:"type:varchar(100)" json:"name_th"`
	NameEN     string    `gorm:"type:varchar(100)" json:"name_en"`
	ProvinceID uuid.UUID `gorm:"type:uuid" json:"province_id"`

	Province *Province `gorm:"foreignKey:ProvinceID"`
}

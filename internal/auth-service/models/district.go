package models

import (
	"github.com/google/uuid"
	"github.com/wisaitas/rbac-golang/pkg"
)

type District struct {
	pkg.BaseModel
	NameTH     string    `gorm:"type:varchar(100);not null"`
	NameEN     string    `gorm:"type:varchar(100);not null"`
	ProvinceID uuid.UUID `gorm:"type:uuid;not null"`

	Province *Province `gorm:"foreignKey:ProvinceID"`
}

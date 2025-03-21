package models

import "github.com/wisaitas/rbac-golang/pkg"

type Province struct {
	pkg.BaseModel
	NameTH string `gorm:"type:varchar(100);not null"`
	NameEN string `gorm:"type:varchar(100);not null"`

	Districts []District `gorm:"foreignKey:ProvinceID"`
}

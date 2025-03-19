package models

import "github.com/wisaitas/rbac-golang/pkg"

type Province struct {
	pkg.BaseModel
	NameTH string `gorm:"type:varchar(100)" json:"name_th"`
	NameEN string `gorm:"type:varchar(100)" json:"name_en"`
}

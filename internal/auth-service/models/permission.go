package models

import "github.com/wisaitas/rbac-golang/pkg"

type Permission struct {
	pkg.BaseModel
	Name        string  `gorm:"type:varchar(255);not null;unique"`
	Description *string `gorm:"type:varchar(400)"`

	Roles []Role `gorm:"many2many:role_permissions;"`
}

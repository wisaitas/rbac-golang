package models

import "github.com/wisaitas/rbac-golang/pkg"

type Role struct {
	pkg.BaseModel
	Name        string  `gorm:"type:varchar(255);not null;unique"`
	Description *string `gorm:"type:varchar(400)"`

	Permissions []Permission `gorm:"many2many:role_permissions;"`
	Users       []User       `gorm:"many2many:user_roles;"`
}

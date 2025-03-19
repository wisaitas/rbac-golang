package models

import (
	"github.com/google/uuid"
	"github.com/wisaitas/rbac-golang/pkg"
)

type UserRole struct {
	pkg.BaseModel
	RoleID uuid.UUID `gorm:"type:uuid;not null"`
	UserID uuid.UUID `gorm:"type:uuid;not null"`

	Role *Role `gorm:"foreignKey:RoleID"`
	User *User `gorm:"foreignKey:UserID"`
}

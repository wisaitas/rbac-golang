package models

import (
	"github.com/google/uuid"
	"github.com/wisaitas/rbac-golang/pkg"
)

type RolePermission struct {
	pkg.BaseModel
	RoleID       uuid.UUID `gorm:"type:uuid;not null"`
	PermissionID uuid.UUID `gorm:"type:uuid;not null"`

	Role       *Role       `gorm:"foreignKey:RoleID"`
	Permission *Permission `gorm:"foreignKey:PermissionID"`
}

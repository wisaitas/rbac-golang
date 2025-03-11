package models

import "github.com/google/uuid"

type RolePermission struct {
	BaseModel
	RoleID       uuid.UUID `gorm:"type:uuid;not null"`
	PermissionID uuid.UUID `gorm:"type:uuid;not null"`

	Role       *Role       `gorm:"foreignKey:RoleID"`
	Permission *Permission `gorm:"foreignKey:PermissionID"`
}

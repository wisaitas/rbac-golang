package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/wisaitas/rbac-golang/pkg"
)

type UserHistory struct {
	pkg.BaseModel
	Action       string    `gorm:"type:varchar(20);not null"`
	UserID       uuid.UUID `gorm:"type:uuid;not null"`
	OldFirstName string    `gorm:"type:varchar(80);"`
	OldLastName  string    `gorm:"type:varchar(80);"`
	OldBirthDate time.Time `gorm:"type:date;"`
	OldPassword  string    `gorm:"type:varchar(255);"`
	OldVersion   int       `gorm:"type:integer;"`
}

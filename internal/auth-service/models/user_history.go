package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/wisaitas/rbac-golang/pkg"
)

type UserHistory struct {
	pkg.BaseModel
	Action      string    `gorm:"type:varchar(20);not null"`
	UserID      uuid.UUID `gorm:"type:uuid;not null"`
	UserVersion int       `gorm:"type:integer;not null"`
	FirstName   string    `gorm:"type:varchar(80);not null"`
	LastName    string    `gorm:"type:varchar(80);not null"`
	BirthDate   time.Time `gorm:"type:date;not null"`
}

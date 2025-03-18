package models

import (
	"time"
)

type User struct {
	BaseModel
	Username  string    `gorm:"type:varchar(40);not null;unique"`
	FirstName string    `gorm:"type:varchar(80);not null"`
	LastName  string    `gorm:"type:varchar(80);not null"`
	BirthDate time.Time `gorm:"type:date;not null"`
	Email     string    `gorm:"type:varchar(255);not null;unique"`
	Password  string    `gorm:"type:varchar(255);not null"`

	Addresses []Address `gorm:"foreignKey:UserID"`
	Roles     []Role    `gorm:"many2many:user_roles;"`
}

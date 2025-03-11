package models

type Permission struct {
	BaseModel
	Name        string  `gorm:"type:varchar(255);not null"`
	Description *string `gorm:"type:varchar(400)"`

	Roles []Role `gorm:"many2many:role_permissions;"`
}

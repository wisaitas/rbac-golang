package models

type Role struct {
	BaseModel
	Name        string  `gorm:"type:varchar(255);not null"`
	Description *string `gorm:"type:varchar(400)"`

	Permissions []Permission `gorm:"many2many:role_permissions;"`
	Users       []User       `gorm:"many2many:user_roles;"`
}

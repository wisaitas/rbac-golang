package models

type Permission struct {
	ID          int     `gorm:"type:integer;primaryKey"`
	Name        string  `gorm:"type:varchar(255);not null;unique"`
	Description *string `gorm:"type:varchar(400)"`

	Roles []Role `gorm:"many2many:role_permissions;"`
}

package models

type Permission struct {
	BaseModel
	Name        string  `gorm:"type:varchar(255);not null"`
	Description *string `gorm:"type:varchar(400)"`

	Roles []Role `gorm:"many2many:role_permissions;"`
}

// ตัวอย่างสิทธิ์ที่ควรสร้าง:
// USER_CREATE, USER_READ, USER_UPDATE, USER_DELETE
// ROLE_CREATE, ROLE_READ, ROLE_UPDATE, ROLE_DELETE
// ADDRESS_CREATE, ADDRESS_READ, ADDRESS_UPDATE, ADDRESS_DELETE

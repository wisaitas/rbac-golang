package models

type Permission struct {
	ID          int     `gorm:"type:integer;primaryKey" json:"id"`
	Name        string  `gorm:"type:varchar(255);not null" json:"name"`
	Description *string `gorm:"type:varchar(400)" json:"description"`

	Roles []Role `gorm:"many2many:role_permissions;" json:"roles"`
}

// ตัวอย่างสิทธิ์ที่ควรสร้าง:
// USER_CREATE, USER_READ, USER_UPDATE, USER_DELETE
// ROLE_CREATE, ROLE_READ, ROLE_UPDATE, ROLE_DELETE
// ADDRESS_CREATE, ADDRESS_READ, ADDRESS_UPDATE, ADDRESS_DELETE

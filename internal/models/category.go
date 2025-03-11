package models

type Category struct {
	BaseModel
	Name        string `gorm:"type:varchar(100);not null"`
	Description string `gorm:"type:text"`

	Products []Product `gorm:"foreignKey:CategoryID"`
}

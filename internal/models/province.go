package models

type Province struct {
	ID     int    `gorm:"type:integer;primaryKey" json:"id"`
	NameTH string `gorm:"type:varchar(100)" json:"name_th"`
	NameEN string `gorm:"type:varchar(100)" json:"name_en"`
}

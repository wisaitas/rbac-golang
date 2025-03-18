package models

type SubDistrict struct {
	ID         int    `gorm:"type:integer;primaryKey" json:"id"`
	NameTH     string `gorm:"type:varchar(100)" json:"name_th"`
	NameEN     string `gorm:"type:varchar(100)" json:"name_en"`
	DistrictID int    `gorm:"type:integer" json:"district_id"`
	ZipCode    int    `gorm:"type:integer" json:"zip_code"`

	District *District `gorm:"foreignKey:DistrictID"`
}

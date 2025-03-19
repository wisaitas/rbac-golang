package district

import (
	"github.com/wisaitas/rbac-golang/internal/auth-service/models"
)

type districtData struct {
	ID         int    `json:"id"`
	NameTH     string `json:"name_th"`
	NameEN     string `json:"name_en"`
	ProvinceID int    `json:"province_id"`
}

func (r *districtData) ToModel() models.District {
	return models.District{
		ImportID: &r.ID,
		NameTH:   r.NameTH,
		NameEN:   r.NameEN,
	}
}

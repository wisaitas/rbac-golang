package province

import "github.com/wisaitas/rbac-golang/internal/auth-service/models"

type provinceData struct {
	ID     int    `json:"id"`
	NameTH string `json:"name_th"`
	NameEN string `json:"name_en"`
}

func (r *provinceData) ToModel() models.Province {
	return models.Province{
		ImportID: &r.ID,
		NameTH:   r.NameTH,
		NameEN:   r.NameEN,
	}
}

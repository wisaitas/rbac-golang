package sub_district

import (
	"strconv"

	"github.com/wisaitas/rbac-golang/internal/auth-service/models"
)

type subDistrictData struct {
	ID         int    `json:"id"`
	NameTH     string `json:"name_th"`
	NameEN     string `json:"name_en"`
	ZipCode    int    `json:"zip_code"`
	DistrictID int    `json:"district_id"`
}

func (r *subDistrictData) ToModel() models.SubDistrict {
	return models.SubDistrict{
		ImportID: r.ID,
		NameTH:   r.NameTH,
		NameEN:   r.NameEN,
		ZipCode:  strconv.Itoa(r.ZipCode),
	}
}

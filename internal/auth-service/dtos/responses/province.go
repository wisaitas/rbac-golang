package responses

import (
	"github.com/wisaitas/rbac-golang/internal/auth-service/models"
	"github.com/wisaitas/rbac-golang/pkg"
)

type ProvinceResponse struct {
	pkg.BaseResponse
	NameTH string `json:"name_th"`
	NameEN string `json:"name_en"`
}

func (r *ProvinceResponse) ModelToResponse(province models.Province) ProvinceResponse {
	r.ID = province.ID
	r.NameTH = province.NameTH
	r.NameEN = province.NameEN

	return *r
}

package responses

import (
	"github.com/google/uuid"
	"github.com/wisaitas/rbac-golang/internal/auth-service/models"
	"github.com/wisaitas/rbac-golang/pkg"
)

type DistrictResponse struct {
	pkg.BaseResponse
	NameTH     string    `json:"name_th"`
	NameEN     string    `json:"name_en"`
	ProvinceID uuid.UUID `json:"province_id"`
}

func (r *DistrictResponse) ModelToResponse(model models.District) DistrictResponse {
	r.ID = model.ID
	r.NameTH = model.NameTH
	r.NameEN = model.NameEN
	r.ProvinceID = model.ProvinceID

	return *r
}

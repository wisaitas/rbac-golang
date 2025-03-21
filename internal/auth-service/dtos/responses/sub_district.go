package responses

import (
	"github.com/google/uuid"
	"github.com/wisaitas/rbac-golang/internal/auth-service/models"
	"github.com/wisaitas/rbac-golang/pkg"
)

type SubDistrictResponse struct {
	pkg.BaseResponse
	NameTH     string    `json:"name_th"`
	NameEN     string    `json:"name_en"`
	DistrictID uuid.UUID `json:"district_id"`
	PostalCode string    `json:"postal_code"`
}

func (r *SubDistrictResponse) ModelToResponse(model models.SubDistrict) SubDistrictResponse {
	r.ID = model.ID
	r.NameTH = model.NameTH
	r.NameEN = model.NameEN
	r.DistrictID = model.DistrictID
	r.PostalCode = model.PostalCode

	return *r
}

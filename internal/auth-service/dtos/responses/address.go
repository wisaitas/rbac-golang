package responses

import (
	"github.com/google/uuid"
	"github.com/wisaitas/rbac-golang/internal/auth-service/models"
	"github.com/wisaitas/rbac-golang/pkg"
)

type AddressResponse struct {
	pkg.BaseResponse
	ProvinceID    uuid.UUID `json:"province_id"`
	DistrictID    uuid.UUID `json:"district_id"`
	SubDistrictID uuid.UUID `json:"sub_district_id"`
	Address       string    `json:"address"`
}

func (r *AddressResponse) ModelToResponse(address models.Address) AddressResponse {
	r.ID = address.ID
	r.CreatedAt = address.CreatedAt
	r.UpdatedAt = address.UpdatedAt
	r.ProvinceID = address.ProvinceID
	r.DistrictID = address.DistrictID
	r.SubDistrictID = address.SubDistrictID
	r.Address = *address.Address

	return *r
}

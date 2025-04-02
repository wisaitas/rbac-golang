package queries

import (
	"github.com/google/uuid"
	"github.com/wisaitas/rbac-golang/pkg"
)

type DistrictQuery struct {
	pkg.PaginationQuery
	ProvinceID uuid.UUID `query:"province_id" validate:"required"`
}

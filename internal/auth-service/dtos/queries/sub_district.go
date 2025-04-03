package queries

import (
	"github.com/google/uuid"
	"github.com/wisaitas/rbac-golang/pkg"
)

type SubDistrictQuery struct {
	pkg.PaginationQuery
	DistrictID uuid.UUID `query:"district_id" validate:"required"`
}

package queries

import "github.com/wisaitas/rbac-golang/pkg"

type SubDistrictQuery struct {
	pkg.PaginationQuery
	DistrictID int `query:"district_id"`
}

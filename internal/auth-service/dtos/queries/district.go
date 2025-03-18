package queries

import "github.com/wisaitas/rbac-golang/pkg"

type DistrictQuery struct {
	pkg.PaginationQuery
	ProvinceID int `query:"province_id"`
}

package queries

import "github.com/wisaitas/rbac-golang/pkg"

type RoleQuery struct {
	pkg.PaginationQuery
	Name string `query:"name"`
}

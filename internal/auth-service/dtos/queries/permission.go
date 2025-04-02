package queries

import "github.com/wisaitas/rbac-golang/pkg"

type PermissionQuery struct {
	pkg.PaginationQuery
	Name *string `query:"name"`
}

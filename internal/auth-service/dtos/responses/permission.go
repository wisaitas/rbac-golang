package responses

import (
	"github.com/wisaitas/rbac-golang/internal/auth-service/models"
	"github.com/wisaitas/rbac-golang/pkg"
)

type PermissionResponse struct {
	pkg.BaseResponse
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

func (r *PermissionResponse) ModelToResponse(permission models.Permission) PermissionResponse {
	r.ID = permission.ID
	r.Name = permission.Name
	r.Description = permission.Description

	return *r
}

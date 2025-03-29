package responses

import (
	"github.com/wisaitas/rbac-golang/internal/auth-service/models"
	"github.com/wisaitas/rbac-golang/pkg"
)

type RoleResponse struct {
	pkg.BaseResponse
	Name        string  `json:"name"`
	Description *string `json:"description"`

	Permissions []PermissionResponse `json:"permissions"`
}

func (r *RoleResponse) ModelToResponse(role models.Role) RoleResponse {
	permissions := []PermissionResponse{}

	for _, permission := range role.Permissions {
		permissionResponse := PermissionResponse{}
		permissionResponse = permissionResponse.ModelToResponse(permission)
		permissions = append(permissions, permissionResponse)
	}

	r.ID = role.ID
	r.CreatedAt = role.CreatedAt
	r.UpdatedAt = role.UpdatedAt
	r.Name = role.Name
	r.Description = role.Description
	r.Permissions = permissions

	return *r
}

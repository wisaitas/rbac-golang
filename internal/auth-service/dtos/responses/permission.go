package responses

import "github.com/wisaitas/rbac-golang/internal/auth-service/models"

type CreatePermissionResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

func (r *CreatePermissionResponse) ModelToResponse(permission models.Permission) CreatePermissionResponse {
	r.ID = permission.ID
	r.Name = permission.Name
	r.Description = permission.Description

	return *r
}

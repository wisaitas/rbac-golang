package requests

import "github.com/wisaitas/rbac-golang/internal/auth-service/models"

type CreateRoleRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description" validate:"omitempty"`

	Permissions []string `json:"permissions" validate:"required"`
}

func (r *CreateRoleRequest) ReqToModel() models.Role {
	permissions := []models.Permission{}
	for _, permission := range r.Permissions {
		permissions = append(permissions, models.Permission{
			Name: permission,
		})
	}

	return models.Role{
		Name:        r.Name,
		Description: r.Description,
		Permissions: permissions,
	}
}

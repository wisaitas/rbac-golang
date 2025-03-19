package responses

import (
	"github.com/google/uuid"
	"github.com/wisaitas/rbac-golang/internal/auth-service/models"
)

type CreateRoleResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
}

func (r *CreateRoleResponse) ModelToResponse(role models.Role) CreateRoleResponse {
	r.ID = role.ID
	r.Name = role.Name
	r.Description = role.Description

	return *r
}

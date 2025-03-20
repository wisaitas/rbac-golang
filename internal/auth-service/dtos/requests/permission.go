package requests

import (
	"mime/multipart"

	"github.com/wisaitas/rbac-golang/internal/auth-service/models"
)

type CreatePermissionRequest struct {
	Name        string  `json:"permission_name" validate:"required"`
	Description *string `json:"description" validate:"omitempty"`
}

func (r *CreatePermissionRequest) ReqToModel() models.Permission {
	return models.Permission{
		Name:        r.Name,
		Description: r.Description,
	}
}

type ImportPermission struct {
	File *multipart.FileHeader `form:"file"`
}

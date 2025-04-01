package requests

import (
	"time"

	"github.com/google/uuid"
	"github.com/wisaitas/rbac-golang/internal/auth-service/models"
)

type CreateUserRequest struct {
	Username        string `json:"username" validate:"required,min=3,max=255"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

func (r *CreateUserRequest) ReqToModel() models.User {
	return models.User{
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
	}
}

type UpdateUserRequest struct {
	FirstName *string    `json:"first_name" validate:"omitempty,min=3,max=255"`
	LastName  *string    `json:"last_name" validate:"omitempty,min=3,max=255"`
	BirthDate *time.Time `json:"birth_date" validate:"omitempty"`
	Email     *string    `json:"email" validate:"omitempty,email"`
}

type AssignRoleRequest struct {
	RoleID uuid.UUID `json:"role_id" validate:"required"`
}

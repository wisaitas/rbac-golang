package permission

import (
	"net/http"

	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/requests"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/responses"
	"github.com/wisaitas/rbac-golang/internal/auth-service/repositories"
	"github.com/wisaitas/rbac-golang/pkg"
)

type Create interface {
	CreatePermission(permission requests.CreatePermissionRequest) (resp responses.PermissionResponse, statusCode int, err error)
}

type create struct {
	permissionRepository repositories.PermissionRepository
	redisUtil            pkg.RedisUtil
}

func NewCreate(
	permissionRepository repositories.PermissionRepository,
	redisUtil pkg.RedisUtil,
) Create {
	return &create{
		permissionRepository: permissionRepository,
		redisUtil:            redisUtil,
	}
}

func (s *create) CreatePermission(permission requests.CreatePermissionRequest) (resp responses.PermissionResponse, statusCode int, err error) {
	return responses.PermissionResponse{}, http.StatusCreated, nil
}

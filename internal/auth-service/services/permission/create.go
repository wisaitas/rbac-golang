package permission

import (
	"net/http"

	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/requests"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/responses"
	"github.com/wisaitas/rbac-golang/internal/auth-service/repositories"
	"github.com/wisaitas/rbac-golang/pkg"
)

type Create interface {
	ImportPermission(req requests.ImportPermission) (statusCode int, err error)
}

type create struct {
	permissionRepository repositories.PermissionRepository
	redisClient          pkg.RedisClient
}

func NewCreate(
	permissionRepository repositories.PermissionRepository,
	redisClient pkg.RedisClient,
) Create {
	return &create{
		permissionRepository: permissionRepository,
		redisClient:          redisClient,
	}
}

func (s *create) CreatePermission(permission requests.CreatePermissionRequest) (resp responses.CreatePermissionResponse, statusCode int, err error) {
	return responses.CreatePermissionResponse{}, http.StatusCreated, nil
}

func (s *create) ImportPermission(req requests.ImportPermission) (statusCode int, err error) {
	return http.StatusCreated, nil
}

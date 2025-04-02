package permission

import (
	"net/http"

	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/requests"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/responses"
	"github.com/wisaitas/rbac-golang/internal/auth-service/repositories"
	"github.com/wisaitas/rbac-golang/pkg"
)

type Post interface {
	CreatePermission(permission requests.CreatePermissionRequest) (resp responses.PermissionResponse, statusCode int, err error)
}

type post struct {
	permissionRepository repositories.PermissionRepository
	redisUtil            pkg.RedisUtil
}

func NewPost(
	permissionRepository repositories.PermissionRepository,
	redisUtil pkg.RedisUtil,
) Post {
	return &post{
		permissionRepository: permissionRepository,
		redisUtil:            redisUtil,
	}
}

func (s *post) CreatePermission(permission requests.CreatePermissionRequest) (resp responses.PermissionResponse, statusCode int, err error) {
	return responses.PermissionResponse{}, http.StatusCreated, nil
}

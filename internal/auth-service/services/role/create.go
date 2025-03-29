package role

import (
	"net/http"

	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/requests"
	"github.com/wisaitas/rbac-golang/internal/auth-service/repositories"
	"github.com/wisaitas/rbac-golang/pkg"
)

type Create interface {
	CreateRole(req requests.CreateRoleRequest) (statusCode int, err error)
}

type create struct {
	roleRepository repositories.RoleRepository
	redisUtil      pkg.RedisUtil
}

func NewCreate(
	roleRepository repositories.RoleRepository,
	redisUtil pkg.RedisUtil,
) Create {
	return &create{
		roleRepository: roleRepository,
		redisUtil:      redisUtil,
	}
}

func (r *create) CreateRole(req requests.CreateRoleRequest) (statusCode int, err error) {
	_ = req.ReqToModel()
	return http.StatusOK, nil
}

package role

import (
	"net/http"

	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/requests"
	"github.com/wisaitas/rbac-golang/internal/auth-service/repositories"
	"github.com/wisaitas/rbac-golang/pkg"
)

type Post interface {
	CreateRole(req requests.CreateRoleRequest) (statusCode int, err error)
}

type post struct {
	roleRepository repositories.RoleRepository
	redisUtil      pkg.RedisUtil
}

func NewPost(
	roleRepository repositories.RoleRepository,
	redisUtil pkg.RedisUtil,
) Post {
	return &post{
		roleRepository: roleRepository,
		redisUtil:      redisUtil,
	}
}

func (r *post) CreateRole(req requests.CreateRoleRequest) (statusCode int, err error) {
	_ = req.ReqToModel()
	return http.StatusOK, nil
}

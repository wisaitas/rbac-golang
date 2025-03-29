package role

import (
	"net/http"

	"github.com/wisaitas/rbac-golang/internal/auth-service/repositories"
	"github.com/wisaitas/rbac-golang/pkg"
)

type Update interface {
	RotateRole() (statusCode int, err error)
}

type update struct {
	roleRepository repositories.RoleRepository
	redisUtil      pkg.RedisUtil
}

func NewUpdate(
	roleRepository repositories.RoleRepository,
	redisUtil pkg.RedisUtil,
) Update {
	return &update{
		roleRepository: roleRepository,
		redisUtil:      redisUtil,
	}
}

func (r *update) RotateRole() (statusCode int, err error) {
	return http.StatusOK, nil
}

package role

import (
	"net/http"

	"github.com/wisaitas/rbac-golang/internal/auth-service/repositories"
	"github.com/wisaitas/rbac-golang/pkg"
)

type Update interface {
	SwitchRole() (statusCode int, err error)
}

type update struct {
	roleRepository repositories.RoleRepository
	redisClient    *pkg.RedisClient
}

func NewUpdate(
	roleRepository repositories.RoleRepository,
	redisClient *pkg.RedisClient,
) Update {
	return &update{
		roleRepository: roleRepository,
		redisClient:    redisClient,
	}
}

func (r *update) SwitchRole() (statusCode int, err error) {
	return http.StatusOK, nil
}

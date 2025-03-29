package role

import (
	"net/http"

	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/queries"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/responses"
	"github.com/wisaitas/rbac-golang/internal/auth-service/models"
	"github.com/wisaitas/rbac-golang/internal/auth-service/repositories"
	"github.com/wisaitas/rbac-golang/pkg"
)

type Read interface {
	GetRoles(query queries.RoleQuery) (resp []responses.RoleResponse, statusCode int, err error)
}

type read struct {
	roleRepository repositories.RoleRepository
	redisUtil      pkg.RedisUtil
}

func NewRead(
	roleRepository repositories.RoleRepository,
	redisUtil pkg.RedisUtil,
) Read {
	return &read{
		roleRepository: roleRepository,
		redisUtil:      redisUtil,
	}
}

func (r *read) GetRoles(query queries.RoleQuery) (resp []responses.RoleResponse, statusCode int, err error) {
	roles := []models.Role{}

	if err := r.roleRepository.GetAll(&roles, nil, nil); err != nil {
		return []responses.RoleResponse{}, http.StatusInternalServerError, err
	}

	for _, role := range roles {
		response := responses.RoleResponse{}
		resp = append(resp, response.ModelToResponse(role))
	}

	return resp, http.StatusOK, nil
}

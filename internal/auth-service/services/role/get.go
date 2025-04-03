package role

import (
	"net/http"

	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/queries"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/responses"
	"github.com/wisaitas/rbac-golang/internal/auth-service/models"
	"github.com/wisaitas/rbac-golang/internal/auth-service/repositories"
	"github.com/wisaitas/rbac-golang/pkg"
)

type Get interface {
	GetRoles(query queries.RoleQuery) (resp []responses.RoleResponse, statusCode int, err error)
}

type get struct {
	roleRepository repositories.RoleRepository
	redisUtil      pkg.RedisUtil
}

func NewGet(
	roleRepository repositories.RoleRepository,
	redisUtil pkg.RedisUtil,
) Get {
	return &get{
		roleRepository: roleRepository,
		redisUtil:      redisUtil,
	}
}

func (r *get) GetRoles(query queries.RoleQuery) (resp []responses.RoleResponse, statusCode int, err error) {
	roles := []models.Role{}

	if err := r.roleRepository.GetAll(&roles, &query.PaginationQuery, nil, nil); err != nil {
		return []responses.RoleResponse{}, http.StatusInternalServerError, err
	}

	for _, role := range roles {
		response := responses.RoleResponse{}
		resp = append(resp, response.ModelToResponse(role))
	}

	return resp, http.StatusOK, nil
}

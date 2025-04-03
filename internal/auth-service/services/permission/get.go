package permission

import (
	"net/http"

	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/queries"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/responses"
	"github.com/wisaitas/rbac-golang/internal/auth-service/models"
	"github.com/wisaitas/rbac-golang/internal/auth-service/repositories"
	"github.com/wisaitas/rbac-golang/pkg"
)

type Get interface {
	GetPermissions(query queries.PermissionQuery) (resp []responses.PermissionResponse, statusCode int, err error)
}

type get struct {
	permissionRepository repositories.PermissionRepository
	redisUtil            pkg.RedisUtil
}

func NewGet(
	permissionRepository repositories.PermissionRepository,
	redisUtil pkg.RedisUtil,
) Get {
	return &get{
		permissionRepository: permissionRepository,
		redisUtil:            redisUtil,
	}
}

func (r *get) GetPermissions(query queries.PermissionQuery) (resp []responses.PermissionResponse, statusCode int, err error) {
	permissions := []models.Permission{}

	var condition *pkg.Condition
	if query.Name != nil {
		condition = pkg.NewCondition("name LIKE ?", "%"+*query.Name+"%")
	}

	if err := r.permissionRepository.GetAll(&permissions, &query.PaginationQuery, condition, nil); err != nil {
		return []responses.PermissionResponse{}, http.StatusInternalServerError, pkg.Error(err)
	}

	for _, permission := range permissions {
		permissionResponse := responses.PermissionResponse{}
		resp = append(resp, permissionResponse.ModelToResponse(permission))
	}

	return resp, http.StatusOK, nil
}

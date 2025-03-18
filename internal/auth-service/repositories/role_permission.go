package repositories

import (
	"github.com/wisaitas/rbac-golang/internal/auth-service/models"
	"github.com/wisaitas/rbac-golang/pkg"

	"gorm.io/gorm"
)

type RolePermissionRepository interface {
	pkg.BaseRepository[models.RolePermission]
}

type rolePermissionRepository struct {
	pkg.BaseRepository[models.RolePermission]
	db *gorm.DB
}

func NewRolePermissionRepository(db *gorm.DB, baseRepository pkg.BaseRepository[models.RolePermission]) RolePermissionRepository {
	return &rolePermissionRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}

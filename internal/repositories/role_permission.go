package repositories

import (
	"github.com/wisaitas/standard-golang/internal/models"

	"gorm.io/gorm"
)

type RolePermissionRepository interface {
	BaseRepository[models.RolePermission]
}

type rolePermissionRepository struct {
	BaseRepository[models.RolePermission]
	db *gorm.DB
}

func NewRolePermissionRepository(db *gorm.DB, baseRepository BaseRepository[models.RolePermission]) RolePermissionRepository {
	return &rolePermissionRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}

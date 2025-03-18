package repositories

import (
	"github.com/wisaitas/rbac-golang/internal/auth-service/models"
	"github.com/wisaitas/rbac-golang/pkg"

	"gorm.io/gorm"
)

type PermissionRepository interface {
	pkg.BaseRepository[models.Permission]
}

type permissionRepository struct {
	pkg.BaseRepository[models.Permission]
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB, baseRepository pkg.BaseRepository[models.Permission]) PermissionRepository {
	return &permissionRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}

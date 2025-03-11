package repositories

import (
	"github.com/wisaitas/standard-golang/internal/models"

	"gorm.io/gorm"
)

type PermissionRepository interface {
	BaseRepository[models.Permission]
}

type permissionRepository struct {
	BaseRepository[models.Permission]
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB, baseRepository BaseRepository[models.Permission]) PermissionRepository {
	return &permissionRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}

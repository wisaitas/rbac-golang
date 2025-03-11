package repositories

import (
	"github.com/wisaitas/standard-golang/internal/models"

	"gorm.io/gorm"
)

type UserRoleRepository interface {
	BaseRepository[models.UserRole]
}

type userRoleRepository struct {
	BaseRepository[models.UserRole]
	db *gorm.DB
}

func NewUserRoleRepository(db *gorm.DB, baseRepository BaseRepository[models.UserRole]) UserRoleRepository {
	return &userRoleRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}

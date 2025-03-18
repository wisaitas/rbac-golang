package repositories

import (
	"github.com/wisaitas/rbac-golang/internal/auth-service/models"
	"github.com/wisaitas/rbac-golang/pkg"

	"gorm.io/gorm"
)

type UserRoleRepository interface {
	pkg.BaseRepository[models.UserRole]
}

type userRoleRepository struct {
	pkg.BaseRepository[models.UserRole]
	db *gorm.DB
}

func NewUserRoleRepository(db *gorm.DB, baseRepository pkg.BaseRepository[models.UserRole]) UserRoleRepository {
	return &userRoleRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}

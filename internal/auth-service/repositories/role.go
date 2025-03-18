package repositories

import (
	"github.com/wisaitas/rbac-golang/internal/auth-service/models"
	"github.com/wisaitas/rbac-golang/pkg"

	"gorm.io/gorm"
)

type RoleRepository interface {
	pkg.BaseRepository[models.Role]
}

type roleRepository struct {
	pkg.BaseRepository[models.Role]
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB, baseRepository pkg.BaseRepository[models.Role]) RoleRepository {
	return &roleRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}

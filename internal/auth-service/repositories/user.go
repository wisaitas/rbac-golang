package repositories

import (
	"github.com/wisaitas/rbac-golang/internal/auth-service/models"
	"github.com/wisaitas/rbac-golang/pkg"

	"gorm.io/gorm"
)

type UserRepository interface {
	pkg.BaseRepository[models.User]
}

type userRepository struct {
	pkg.BaseRepository[models.User]
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB, baseRepository pkg.BaseRepository[models.User]) UserRepository {
	return &userRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}

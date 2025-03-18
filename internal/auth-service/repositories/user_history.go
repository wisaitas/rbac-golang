package repositories

import (
	"github.com/wisaitas/rbac-golang/internal/auth-service/models"
	"github.com/wisaitas/rbac-golang/pkg"
	"gorm.io/gorm"
)

type UserHistoryRepository interface {
	pkg.BaseRepository[models.UserHistory]
}

type userHistoryRepository struct {
	pkg.BaseRepository[models.UserHistory]
	db *gorm.DB
}

func NewUserHistoryRepository(db *gorm.DB, baseRepository pkg.BaseRepository[models.UserHistory]) UserHistoryRepository {
	return &userHistoryRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}

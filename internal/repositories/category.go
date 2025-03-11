package repositories

import (
	"github.com/wisaitas/standard-golang/internal/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	BaseRepository[models.Category]
}

type categoryRepository struct {
	BaseRepository[models.Category]
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB, baseRepository BaseRepository[models.Category]) CategoryRepository {
	return &categoryRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}

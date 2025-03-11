package repositories

import (
	"github.com/wisaitas/standard-golang/internal/models"
	"gorm.io/gorm"
)

type ProductImageRepository interface {
	BaseRepository[models.ProductImage]
}

type productImageRepository struct {
	BaseRepository[models.ProductImage]
	db *gorm.DB
}

func NewProductImageRepository(db *gorm.DB, baseRepository BaseRepository[models.ProductImage]) ProductImageRepository {
	return &productImageRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}

package repositories

import (
	"github.com/wisaitas/standard-golang/internal/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	BaseRepository[models.Product]
}

type productRepository struct {
	BaseRepository[models.Product]
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB, baseRepository BaseRepository[models.Product]) ProductRepository {
	return &productRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}

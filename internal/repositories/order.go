package repositories

import (
	"github.com/wisaitas/standard-golang/internal/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	BaseRepository[models.Order]
}

type orderRepository struct {
	BaseRepository[models.Order]
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB, baseRepository BaseRepository[models.Order]) OrderRepository {
	return &orderRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}

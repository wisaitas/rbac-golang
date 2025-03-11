package repositories

import (
	"github.com/wisaitas/standard-golang/internal/models"
	"gorm.io/gorm"
)

type OrderItemRepository interface {
	BaseRepository[models.OrderItem]
}

type orderItemRepository struct {
	BaseRepository[models.OrderItem]
	db *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB, baseRepository BaseRepository[models.OrderItem]) OrderItemRepository {
	return &orderItemRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}

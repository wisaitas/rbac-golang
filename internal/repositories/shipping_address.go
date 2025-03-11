package repositories

import (
	"github.com/wisaitas/standard-golang/internal/models"
	"gorm.io/gorm"
)

type ShippingAddressRepository interface {
	BaseRepository[models.ShippingAddress]
}

type shippingAddressRepository struct {
	BaseRepository[models.ShippingAddress]
	db *gorm.DB
}

func NewShippingAddressRepository(db *gorm.DB, baseRepository BaseRepository[models.ShippingAddress]) ShippingAddressRepository {
	return &shippingAddressRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}

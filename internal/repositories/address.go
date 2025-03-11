package repositories

import (
	"github.com/wisaitas/standard-golang/internal/models"

	"gorm.io/gorm"
)

type AddressRepository interface {
	BaseRepository[models.Address]
}

type addressRepository struct {
	BaseRepository[models.Address]
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB, baseRepository BaseRepository[models.Address]) AddressRepository {
	return &addressRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}

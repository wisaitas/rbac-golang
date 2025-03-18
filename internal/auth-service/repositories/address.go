package repositories

import (
	"github.com/wisaitas/rbac-golang/internal/auth-service/models"
	"github.com/wisaitas/rbac-golang/pkg"

	"gorm.io/gorm"
)

type AddressRepository interface {
	pkg.BaseRepository[models.Address]
}

type addressRepository struct {
	pkg.BaseRepository[models.Address]
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB, baseRepository pkg.BaseRepository[models.Address]) AddressRepository {
	return &addressRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}

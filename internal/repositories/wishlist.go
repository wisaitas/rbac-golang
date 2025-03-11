package repositories

import (
	"github.com/wisaitas/standard-golang/internal/models"
	"gorm.io/gorm"
)

type WishlistRepository interface {
	BaseRepository[models.Wishlist]
}

type wishlistRepository struct {
	BaseRepository[models.Wishlist]
	db *gorm.DB
}

func NewWishlistRepository(db *gorm.DB, baseRepository BaseRepository[models.Wishlist]) WishlistRepository {
	return &wishlistRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}

package repositories

import (
	"github.com/wisaitas/rbac-golang/internal/auth-service/models"
	"github.com/wisaitas/rbac-golang/pkg"
	"gorm.io/gorm"
)

type ProvinceRepository interface {
	pkg.BaseRepository[models.Province]
}

type provinceRepository struct {
	pkg.BaseRepository[models.Province]
	db *gorm.DB
}

func NewProvinceRepository(db *gorm.DB, baseRepository pkg.BaseRepository[models.Province]) ProvinceRepository {
	return &provinceRepository{
		BaseRepository: baseRepository,
		db:             db,
	}
}

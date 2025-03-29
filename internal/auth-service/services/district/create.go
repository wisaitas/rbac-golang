package district

import (
	"github.com/wisaitas/rbac-golang/internal/auth-service/repositories"
	"github.com/wisaitas/rbac-golang/pkg"
)

type Create interface {
}

type create struct {
	districtRepository repositories.DistrictRepository
	redisUtil          pkg.RedisUtil
}

func NewCreate(
	districtRepository repositories.DistrictRepository,
	redisUtil pkg.RedisUtil,
) Create {
	return &create{
		districtRepository: districtRepository,
		redisUtil:          redisUtil,
	}
}

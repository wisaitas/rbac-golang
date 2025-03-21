package district

import (
	"github.com/wisaitas/rbac-golang/internal/auth-service/repositories"
	"github.com/wisaitas/rbac-golang/pkg"
)

type Create interface {
}

type create struct {
	districtRepository repositories.DistrictRepository
	redisClient        pkg.RedisClient
}

func NewCreate(
	districtRepository repositories.DistrictRepository,
	redisClient pkg.RedisClient,
) Create {
	return &create{
		districtRepository: districtRepository,
		redisClient:        redisClient,
	}
}

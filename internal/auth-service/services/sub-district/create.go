package sub_district

import (
	"github.com/wisaitas/rbac-golang/internal/auth-service/repositories"
	"github.com/wisaitas/rbac-golang/pkg"
)

type Create interface {
}

type create struct {
	subDistrictRepository repositories.SubDistrictRepository
	redisClient           pkg.RedisClient
}

func NewCreate(
	subDistrictRepository repositories.SubDistrictRepository,
	redisClient pkg.RedisClient,
) Create {
	return &create{
		subDistrictRepository: subDistrictRepository,
		redisClient:           redisClient,
	}
}

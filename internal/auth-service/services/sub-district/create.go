package sub_district

import (
	"github.com/wisaitas/rbac-golang/internal/auth-service/repositories"
	"github.com/wisaitas/rbac-golang/pkg"
)

type Create interface {
}

type create struct {
	subDistrictRepository repositories.SubDistrictRepository
	redisUtil             pkg.RedisUtil
}

func NewCreate(
	subDistrictRepository repositories.SubDistrictRepository,
	redisUtil pkg.RedisUtil,
) Create {
	return &create{
		subDistrictRepository: subDistrictRepository,
		redisUtil:             redisUtil,
	}
}

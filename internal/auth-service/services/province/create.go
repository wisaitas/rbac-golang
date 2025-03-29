package province

import (
	"github.com/wisaitas/rbac-golang/internal/auth-service/repositories"
	"github.com/wisaitas/rbac-golang/pkg"
)

type Create interface {
}

type create struct {
	provinceRepository repositories.ProvinceRepository
	redisUtil          pkg.RedisUtil
}

func NewCreate(
	provinceRepository repositories.ProvinceRepository,
	redisUtil pkg.RedisUtil,
) Create {
	return &create{
		provinceRepository: provinceRepository,
		redisUtil:          redisUtil,
	}
}

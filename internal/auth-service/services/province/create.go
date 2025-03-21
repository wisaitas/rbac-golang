package province

import (
	"github.com/wisaitas/rbac-golang/internal/auth-service/repositories"
	"github.com/wisaitas/rbac-golang/pkg"
)

type Create interface {
}

type create struct {
	provinceRepository repositories.ProvinceRepository
	redisClient        pkg.RedisClient
}

func NewCreate(
	provinceRepository repositories.ProvinceRepository,
	redisClient pkg.RedisClient,
) Create {
	return &create{
		provinceRepository: provinceRepository,
		redisClient:        redisClient,
	}
}

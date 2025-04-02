package district

import (
	"github.com/wisaitas/rbac-golang/internal/auth-service/repositories"
	"github.com/wisaitas/rbac-golang/pkg"
)

type Post interface {
}

type post struct {
	districtRepository repositories.DistrictRepository
	redisUtil          pkg.RedisUtil
}

func NewPost(
	districtRepository repositories.DistrictRepository,
	redisUtil pkg.RedisUtil,
) Post {
	return &post{
		districtRepository: districtRepository,
		redisUtil:          redisUtil,
	}
}

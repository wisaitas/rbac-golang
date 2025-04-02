package sub_district

import (
	"github.com/wisaitas/rbac-golang/internal/auth-service/repositories"
	"github.com/wisaitas/rbac-golang/pkg"
)

type Post interface {
}

type post struct {
	subDistrictRepository repositories.SubDistrictRepository
	redisUtil             pkg.RedisUtil
}

func NewPost(
	subDistrictRepository repositories.SubDistrictRepository,
	redisUtil pkg.RedisUtil,
) Post {
	return &post{
		subDistrictRepository: subDistrictRepository,
		redisUtil:             redisUtil,
	}
}

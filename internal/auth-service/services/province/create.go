package province

import (
	"github.com/wisaitas/rbac-golang/internal/auth-service/repositories"
	"github.com/wisaitas/rbac-golang/pkg"
)

type Post interface {
}

type post struct {
	provinceRepository repositories.ProvinceRepository
	redisUtil          pkg.RedisUtil
}

func NewPost(
	provinceRepository repositories.ProvinceRepository,
	redisUtil pkg.RedisUtil,
) Post {
	return &post{
		provinceRepository: provinceRepository,
		redisUtil:          redisUtil,
	}
}

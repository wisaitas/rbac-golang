package user

import (
	"github.com/wisaitas/rbac-golang/internal/auth-service/repositories"
	"github.com/wisaitas/rbac-golang/pkg"
)

type Delete interface {
}

type delete struct {
	userRepository repositories.UserRepository
	redisUtil      pkg.RedisUtil
}

func NewDelete(
	userRepository repositories.UserRepository,
	redisUtil pkg.RedisUtil,
) Delete {
	return &delete{
		userRepository: userRepository,
		redisUtil:      redisUtil,
	}
}

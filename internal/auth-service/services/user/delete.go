package user

import (
	"github.com/wisaitas/rbac-golang/internal/auth-service/repositories"
	"github.com/wisaitas/rbac-golang/pkg"
)

type Delete interface {
}

type delete struct {
	userRepository repositories.UserRepository
	redisUtil      pkg.RedisClient
}

func NewDelete(
	userRepository repositories.UserRepository,
	redisUtil pkg.RedisClient,
) Delete {
	return &delete{
		userRepository: userRepository,
		redisUtil:      redisUtil,
	}
}

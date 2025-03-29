package initial

import (
	"github.com/wisaitas/rbac-golang/internal/auth-service/middlewares"
)

type Middlewares struct {
	AuthMiddleware middlewares.AuthMiddleware
	UserMiddleware middlewares.UserMiddleware
}

func initializeMiddlewares(utils *Utils) *Middlewares {
	return &Middlewares{
		AuthMiddleware: *middlewares.NewAuthMiddleware(utils.RedisUtil, utils.JWTUtil),
		UserMiddleware: *middlewares.NewUserMiddleware(utils.RedisUtil, utils.JWTUtil),
	}
}

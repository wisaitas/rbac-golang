package initial

import (
	"github.com/wisaitas/rbac-golang/internal/auth-service/middlewares"
	"github.com/wisaitas/rbac-golang/pkg"
)

type Middlewares struct {
	AuthMiddleware middlewares.AuthMiddleware
	UserMiddleware middlewares.UserMiddleware
}

func initializeMiddlewares(redis pkg.RedisClient) *Middlewares {
	return &Middlewares{
		AuthMiddleware: *middlewares.NewAuthMiddleware(redis),
		UserMiddleware: *middlewares.NewUserMiddleware(redis),
	}
}

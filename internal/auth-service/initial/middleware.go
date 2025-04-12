package initial

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/rbac-golang/internal/auth-service/middlewares"
	middlewareConfig "github.com/wisaitas/rbac-golang/internal/auth-service/middlewares/configs"
)

type middleware struct {
	AuthMiddleware middlewares.AuthMiddleware
	UserMiddleware middlewares.UserMiddleware
}

func initializeMiddleware(util *util) *middleware {
	return &middleware{
		AuthMiddleware: *middlewares.NewAuthMiddleware(util.RedisUtil, util.JWTUtil),
		UserMiddleware: *middlewares.NewUserMiddleware(util.RedisUtil, util.JWTUtil),
	}
}

func setupMiddlewares(app *fiber.App) {
	app.Use(
		middlewareConfig.Recovery(),
		middlewareConfig.Limiter(),
		middlewareConfig.CORS(),
		middlewareConfig.Logger(),
		middlewareConfig.Healthz(),
		middlewareConfig.Pprof(),
	)
}

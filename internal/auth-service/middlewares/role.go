package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/rbac-golang/pkg"
)

type RoleMiddleware struct {
	redisUtil pkg.RedisUtil
	jwtUtil   pkg.JWTUtil
}

func NewRoleMiddleware(
	redisUtil pkg.RedisUtil,
	jwtUtil pkg.JWTUtil,
) *RoleMiddleware {
	return &RoleMiddleware{
		redisUtil: redisUtil,
		jwtUtil:   jwtUtil,
	}
}

func (r *RoleMiddleware) RotateRole(c *fiber.Ctx) error {
	return c.Next()
}

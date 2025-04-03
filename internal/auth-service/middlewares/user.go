package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/rbac-golang/pkg"
)

type UserMiddleware struct {
	redisUtil pkg.RedisUtil
	jwtUtil   pkg.JWTUtil
}

func NewUserMiddleware(
	redisUtil pkg.RedisUtil,
	jwtUtil pkg.JWTUtil,
) *UserMiddleware {
	return &UserMiddleware{
		redisUtil: redisUtil,
		jwtUtil:   jwtUtil,
	}
}

func (r *UserMiddleware) GetUserProfile(c *fiber.Ctx) error {
	if err := authToken(c, r.redisUtil, r.jwtUtil); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	return c.Next()
}

func (r *UserMiddleware) UpdateUser(c *fiber.Ctx) error {
	if err := authToken(c, r.redisUtil, r.jwtUtil); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	// if handler permission
	// userContext, ok := c.Locals("userContext").(models.UserContext)
	// if !ok {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(responses.ErrorResponse{
	// 		Message: "user context not found",
	// 	})
	// }

	// if userContext.Username != "test" {
	// 	return c.Status(fiber.StatusForbidden).JSON(responses.ErrorResponse{
	// 		Message: "you are not authorized to access this resource",
	// 	})
	// }

	return c.Next()
}

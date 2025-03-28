package validates

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/requests"
	"github.com/wisaitas/rbac-golang/pkg"
)

type AuthValidate struct {
}

func NewAuthValidate() *AuthValidate {
	return &AuthValidate{}
}

func (r *AuthValidate) ValidateLoginRequest(c *fiber.Ctx) error {
	req := requests.LoginRequest{}

	if err := validateCommonRequestJSONBody(c, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	c.Locals("req", req)
	return c.Next()
}

func (r *AuthValidate) ValidateRegisterRequest(c *fiber.Ctx) error {
	req := requests.RegisterRequest{}

	if err := validateCommonRequestJSONBody(c, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	c.Locals("req", req)
	return c.Next()
}

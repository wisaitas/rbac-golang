package validates

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/requests"
	"github.com/wisaitas/rbac-golang/pkg"
)

type AuthValidate interface {
	ValidateLoginRequest(c *fiber.Ctx) error
	ValidateRegisterRequest(c *fiber.Ctx) error
}

type authValidate struct {
	validatorUtil pkg.ValidatorUtil
}

func NewAuthValidate(validatorUtil pkg.ValidatorUtil) AuthValidate {
	return &authValidate{
		validatorUtil: validatorUtil,
	}
}

func (r *authValidate) ValidateLoginRequest(c *fiber.Ctx) error {
	req := requests.LoginRequest{}

	if err := validateCommonRequestJSONBody(c, &req, r.validatorUtil); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	c.Locals("req", req)
	return c.Next()
}

func (r *authValidate) ValidateRegisterRequest(c *fiber.Ctx) error {
	req := requests.RegisterRequest{}

	if err := validateCommonRequestJSONBody(c, &req, r.validatorUtil); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	c.Locals("req", req)
	return c.Next()
}

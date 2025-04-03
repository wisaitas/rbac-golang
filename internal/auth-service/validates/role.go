package validates

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/queries"
	"github.com/wisaitas/rbac-golang/pkg"
)

type RoleValidate interface {
	ValidateGetRolesRequest(c *fiber.Ctx) error
}

type roleValidate struct {
	validatorUtil pkg.ValidatorUtil
}

func NewRoleValidate(validatorUtil pkg.ValidatorUtil) RoleValidate {
	return &roleValidate{
		validatorUtil: validatorUtil,
	}
}

func (v *roleValidate) ValidateGetRolesRequest(c *fiber.Ctx) error {
	query := queries.RoleQuery{}

	if err := validateCommonPaginationQuery(c, &query, v.validatorUtil); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	c.Locals("query", query)
	return c.Next()
}

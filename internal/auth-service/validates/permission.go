package validates

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/queries"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/requests"
	"github.com/wisaitas/rbac-golang/pkg"
)

type PermissionValidate interface {
	ValidateCreatePermissionRequest(c *fiber.Ctx) error
	ValidateGetPermissionsRequest(c *fiber.Ctx) error
}

type permissionValidate struct {
	validatorUtil pkg.ValidatorUtil
}

func NewPermissionValidate(validatorUtil pkg.ValidatorUtil) PermissionValidate {
	return &permissionValidate{
		validatorUtil: validatorUtil,
	}
}

func (r *permissionValidate) ValidateCreatePermissionRequest(c *fiber.Ctx) error {
	request := requests.CreatePermissionRequest{}

	if err := validateCommonRequestJSONBody(c, &request, r.validatorUtil); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	c.Locals("req", request)
	return c.Next()
}

func (r *permissionValidate) ValidateGetPermissionsRequest(c *fiber.Ctx) error {
	query := queries.PermissionQuery{}

	if err := validateCommonPaginationQuery(c, &query, r.validatorUtil); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	c.Locals("query", query)
	return c.Next()
}

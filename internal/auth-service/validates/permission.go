package validates

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/requests"
	"github.com/wisaitas/rbac-golang/pkg"
)

type PermissionValidate struct {
}

func NewPermissionValidate() *PermissionValidate {
	return &PermissionValidate{}
}

func (r *PermissionValidate) ValidateCreatePermissionRequest(c *fiber.Ctx) error {
	request := requests.CreatePermissionRequest{}

	if err := validateCommonRequestJSONBody(c, &request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	c.Locals("req", request)
	return c.Next()
}

package validates

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/queries"
	"github.com/wisaitas/rbac-golang/pkg"
)

type RoleValidate struct {
}

func NewRoleValidate() *RoleValidate {
	return &RoleValidate{}
}

func (v *RoleValidate) ValidateGetRolesRequest(c *fiber.Ctx) error {
	query := queries.RoleQuery{}

	if err := validateCommonPaginationQuery(c, &query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	c.Locals("query", query)
	return c.Next()
}

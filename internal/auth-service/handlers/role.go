package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/queries"
	roleService "github.com/wisaitas/rbac-golang/internal/auth-service/services/role"
	"github.com/wisaitas/rbac-golang/pkg"
)

type RoleHandler struct {
	roleService roleService.RoleService
}

func NewRoleHandler(roleService roleService.RoleService) *RoleHandler {
	return &RoleHandler{
		roleService: roleService,
	}
}

func (r *RoleHandler) GetRoles(c *fiber.Ctx) error {
	query, ok := c.Locals("query").(queries.RoleQuery)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(errors.New("failed to get queries")).Error(),
		})
	}

	resp, statusCode, err := r.roleService.GetRoles(query)
	if err != nil {
		return c.Status(statusCode).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	return c.Status(statusCode).JSON(resp)
}

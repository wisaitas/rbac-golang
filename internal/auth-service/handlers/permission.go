package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/queries"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/requests"
	permissionService "github.com/wisaitas/rbac-golang/internal/auth-service/services/permission"
	"github.com/wisaitas/rbac-golang/pkg"
)

type PermissionHandler struct {
	permissionService permissionService.PermissionService
}

func NewPermissionHandler(
	permissionService permissionService.PermissionService,
) *PermissionHandler {
	return &PermissionHandler{
		permissionService: permissionService,
	}
}

func (r *PermissionHandler) CreatePermission(c *fiber.Ctx) error {
	req, ok := c.Locals("req").(requests.CreatePermissionRequest)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(errors.New("failed to get request")).Error(),
		})
	}

	resp, statusCode, err := r.permissionService.CreatePermission(req)
	if err != nil {
		return c.Status(statusCode).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	return c.Status(statusCode).JSON(pkg.SuccessResponse{
		Message: "permission created successfully",
		Data:    resp,
	})
}

func (r *PermissionHandler) GetPermissions(c *fiber.Ctx) error {
	query, ok := c.Locals("query").(queries.PermissionQuery)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(errors.New("failed to get query")).Error(),
		})
	}

	resp, statusCode, err := r.permissionService.GetPermissions(query)
	if err != nil {
		return c.Status(statusCode).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	return c.Status(statusCode).JSON(pkg.SuccessResponse{
		Message: "permissions fetched successfully",
		Data:    resp,
	})
}

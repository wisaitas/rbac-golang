package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/rbac-golang/internal/auth-service/handlers"
	"github.com/wisaitas/rbac-golang/internal/auth-service/validates"
)

type PermissionRoutes struct {
	app                fiber.Router
	permissionHandler  *handlers.PermissionHandler
	permissionValidate validates.PermissionValidate
}

func NewPermissionRoutes(
	app fiber.Router,
	permissionHandler *handlers.PermissionHandler,
	permissionValidate validates.PermissionValidate,
) *PermissionRoutes {
	return &PermissionRoutes{
		app:                app,
		permissionHandler:  permissionHandler,
		permissionValidate: permissionValidate,
	}
}

func (r *PermissionRoutes) PermissionRoutes() {
	permissions := r.app.Group("/permissions")

	// Method GET
	permissions.Get("/", r.permissionValidate.ValidateGetPermissionsRequest, r.permissionHandler.GetPermissions)

	// Method POST
	permissions.Post("/", r.permissionValidate.ValidateCreatePermissionRequest, r.permissionHandler.CreatePermission)
}

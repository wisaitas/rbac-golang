package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/rbac-golang/internal/auth-service/handlers"
	"github.com/wisaitas/rbac-golang/internal/auth-service/validates"
)

type RoleRoutes struct {
	app          fiber.Router
	roleHandler  *handlers.RoleHandler
	roleValidate *validates.RoleValidate
}

func NewRoleRoutes(
	app fiber.Router,
	roleHandler *handlers.RoleHandler,
	roleValidate *validates.RoleValidate,
) *RoleRoutes {
	return &RoleRoutes{
		app:          app,
		roleHandler:  roleHandler,
		roleValidate: roleValidate,
	}
}

func (r *RoleRoutes) RoleRoutes() {
	roles := r.app.Group("/roles")

	// Method GET
	roles.Get("/", r.roleValidate.ValidateGetRolesRequest, r.roleHandler.GetRoles)

	// Method Patch
	// roles.Patch("/rotate", r.roleValidate.ValidateRotateRoleRequest, r.roleHandler.RotateRole)
}

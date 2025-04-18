package initial

import (
	"github.com/wisaitas/rbac-golang/internal/auth-service/handlers"
)

type Handlers struct {
	UserHandler        handlers.UserHandler
	AuthHandler        handlers.AuthHandler
	ProvinceHandler    handlers.ProvinceHandler
	DistrictHandler    handlers.DistrictHandler
	SubDistrictHandler handlers.SubDistrictHandler
	PermissionHandler  handlers.PermissionHandler
	RoleHandler        handlers.RoleHandler
}

func initializeHandlers(services *Services) *Handlers {
	return &Handlers{
		UserHandler:        *handlers.NewUserHandler(services.UserService),
		AuthHandler:        *handlers.NewAuthHandler(services.AuthService),
		ProvinceHandler:    *handlers.NewProvinceHandler(services.ProvinceService),
		DistrictHandler:    *handlers.NewDistrictHandler(services.DistrictService),
		SubDistrictHandler: *handlers.NewSubDistrictHandler(services.SubDistrictService),
		PermissionHandler:  *handlers.NewPermissionHandler(services.PermissionService),
		RoleHandler:        *handlers.NewRoleHandler(services.RoleService),
	}
}

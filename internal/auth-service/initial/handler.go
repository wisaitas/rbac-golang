package initial

import (
	"github.com/wisaitas/rbac-golang/internal/auth-service/handlers"
)

type handler struct {
	UserHandler        handlers.UserHandler
	AuthHandler        handlers.AuthHandler
	ProvinceHandler    handlers.ProvinceHandler
	DistrictHandler    handlers.DistrictHandler
	SubDistrictHandler handlers.SubDistrictHandler
	PermissionHandler  handlers.PermissionHandler
	RoleHandler        handlers.RoleHandler
}

func initializeHandler(service *service) *handler {
	return &handler{
		UserHandler:        *handlers.NewUserHandler(service.UserService),
		AuthHandler:        *handlers.NewAuthHandler(service.AuthService),
		ProvinceHandler:    *handlers.NewProvinceHandler(service.ProvinceService),
		DistrictHandler:    *handlers.NewDistrictHandler(service.DistrictService),
		SubDistrictHandler: *handlers.NewSubDistrictHandler(service.SubDistrictService),
		PermissionHandler:  *handlers.NewPermissionHandler(service.PermissionService),
		RoleHandler:        *handlers.NewRoleHandler(service.RoleService),
	}
}

package initial

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/rbac-golang/internal/auth-service/routes"
)

type Routes struct {
	UserRoutes        *routes.UserRoutes
	AuthRoutes        *routes.AuthRoutes
	ProvinceRoutes    *routes.ProvinceRoutes
	DistrictRoutes    *routes.DistrictRoutes
	SubDistrictRoutes *routes.SubDistrictRoutes
	PermissionRoutes  *routes.PermissionRoutes
	RoleRoutes        *routes.RoleRoutes
}

func initializeRoutes(
	apiRoutes fiber.Router,
	handlers *Handlers,
	validates *Validates,
	middlewares *Middlewares,
) *Routes {
	return &Routes{
		UserRoutes: routes.NewUserRoutes(
			apiRoutes,
			&handlers.UserHandler,
			&validates.UserValidate,
			&middlewares.AuthMiddleware,
			&middlewares.UserMiddleware,
		),
		AuthRoutes: routes.NewAuthRoutes(
			apiRoutes,
			&handlers.AuthHandler,
			&validates.AuthValidate,
			&middlewares.AuthMiddleware,
		),
		ProvinceRoutes: routes.NewProvinceRoutes(
			apiRoutes,
			&handlers.ProvinceHandler,
			&validates.ProvinceValidate,
		),
		DistrictRoutes: routes.NewDistrictRoutes(
			apiRoutes,
			&handlers.DistrictHandler,
			&validates.DistrictValidate,
		),
		SubDistrictRoutes: routes.NewSubDistrictRoutes(
			apiRoutes,
			&handlers.SubDistrictHandler,
			&validates.SubDistrictValidate,
		),
		PermissionRoutes: routes.NewPermissionRoutes(
			apiRoutes,
			&handlers.PermissionHandler,
			&validates.PermissionValidate,
		),
		RoleRoutes: routes.NewRoleRoutes(
			apiRoutes,
			&handlers.RoleHandler,
			&validates.RoleValidate,
		),
	}
}

func (r *Routes) SetupRoutes() {
	r.UserRoutes.UserRoutes()
	r.AuthRoutes.AuthRoutes()
	r.ProvinceRoutes.ProvinceRoutes()
	r.DistrictRoutes.DistrictRoutes()
	r.SubDistrictRoutes.SubDistrictRoutes()
	r.PermissionRoutes.PermissionRoutes()
	r.RoleRoutes.RoleRoutes()
}

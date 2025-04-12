package initial

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/rbac-golang/internal/auth-service/routes"
)

type Route struct {
	UserRoutes        *routes.UserRoutes
	AuthRoutes        *routes.AuthRoutes
	ProvinceRoutes    *routes.ProvinceRoutes
	DistrictRoutes    *routes.DistrictRoutes
	SubDistrictRoutes *routes.SubDistrictRoutes
	PermissionRoutes  *routes.PermissionRoutes
	RoleRoutes        *routes.RoleRoutes
}

func initializeRoute(
	app *fiber.App,
	handler *handler,
	validate *validate,
	middleware *middleware,
) {
	apiRoute := app.Group("/api/v1")

	route := &Route{
		UserRoutes: routes.NewUserRoutes(
			apiRoute,
			&handler.UserHandler,
			validate.UserValidate,
			middleware.AuthMiddleware,
			middleware.UserMiddleware,
		),
		AuthRoutes: routes.NewAuthRoutes(
			apiRoute,
			&handler.AuthHandler,
			validate.AuthValidate,
			middleware.AuthMiddleware,
		),
		ProvinceRoutes: routes.NewProvinceRoutes(
			apiRoute,
			&handler.ProvinceHandler,
			validate.ProvinceValidate,
		),
		DistrictRoutes: routes.NewDistrictRoutes(
			apiRoute,
			&handler.DistrictHandler,
			validate.DistrictValidate,
		),
		SubDistrictRoutes: routes.NewSubDistrictRoutes(
			apiRoute,
			&handler.SubDistrictHandler,
			validate.SubDistrictValidate,
		),
		PermissionRoutes: routes.NewPermissionRoutes(
			apiRoute,
			&handler.PermissionHandler,
			validate.PermissionValidate,
		),
		RoleRoutes: routes.NewRoleRoutes(
			apiRoute,
			&handler.RoleHandler,
			validate.RoleValidate,
		),
	}

	route.setupRoute()
}

func (r *Route) setupRoute() {
	r.UserRoutes.UserRoutes()
	r.AuthRoutes.AuthRoutes()
	r.ProvinceRoutes.ProvinceRoutes()
	r.DistrictRoutes.DistrictRoutes()
	r.SubDistrictRoutes.SubDistrictRoutes()
	r.PermissionRoutes.PermissionRoutes()
	r.RoleRoutes.RoleRoutes()
}

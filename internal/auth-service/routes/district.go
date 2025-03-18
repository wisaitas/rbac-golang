package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/rbac-golang/internal/auth-service/handlers"
	"github.com/wisaitas/rbac-golang/internal/auth-service/validates"
)

type DistrictRoutes struct {
	app              fiber.Router
	districtHandler  *handlers.DistrictHandler
	districtValidate *validates.DistrictValidate
}

func NewDistrictRoutes(
	app fiber.Router,
	districtHandler *handlers.DistrictHandler,
	districtValidate *validates.DistrictValidate,
) *DistrictRoutes {
	return &DistrictRoutes{
		app:              app,
		districtHandler:  districtHandler,
		districtValidate: districtValidate,
	}
}

func (r *DistrictRoutes) DistrictRoutes() {
	districts := r.app.Group("/districts")

	// Method GET
	districts.Get("/", r.districtValidate.ValidateGetDistrictsRequest, r.districtHandler.GetDistricts)
}

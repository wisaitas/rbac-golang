package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/queries"
	districtService "github.com/wisaitas/rbac-golang/internal/auth-service/services/district"
	"github.com/wisaitas/rbac-golang/pkg"
)

type DistrictHandler struct {
	districtService districtService.DistrictService
}

func NewDistrictHandler(
	districtService districtService.DistrictService,
) *DistrictHandler {
	return &DistrictHandler{
		districtService: districtService,
	}
}

func (r *DistrictHandler) GetDistricts(c *fiber.Ctx) error {
	query, ok := c.Locals("query").(queries.DistrictQuery)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(errors.New("failed to get queries")).Error(),
		})
	}

	districts, statusCode, err := r.districtService.GetDistricts(query)
	if err != nil {
		return c.Status(statusCode).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	return c.Status(statusCode).JSON(pkg.SuccessResponse{
		Message: "districts fetched successfully",
		Data:    districts,
	})
}

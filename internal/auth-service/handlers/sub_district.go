package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/queries"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/requests"
	subDistrictService "github.com/wisaitas/rbac-golang/internal/auth-service/services/sub-district"
	"github.com/wisaitas/rbac-golang/pkg"
)

type SubDistrictHandler struct {
	subDistrictService subDistrictService.SubDistrictService
}

func NewSubDistrictHandler(
	subDistrictService subDistrictService.SubDistrictService,
) *SubDistrictHandler {
	return &SubDistrictHandler{
		subDistrictService: subDistrictService,
	}
}

func (r *SubDistrictHandler) GetSubDistricts(c *fiber.Ctx) error {
	query, ok := c.Locals("query").(queries.SubDistrictQuery)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(errors.New("failed to get queries")).Error(),
		})
	}

	subDistricts, statusCode, err := r.subDistrictService.GetSubDistricts(query)
	if err != nil {
		return c.Status(statusCode).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	return c.Status(statusCode).JSON(pkg.SuccessResponse{
		Message: "sub-districts fetched successfully",
		Data:    subDistricts,
	})
}

func (r *SubDistrictHandler) ImportSubDistricts(c *fiber.Ctx) error {
	req, ok := c.Locals("req").(requests.ImportSubDistrict)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(errors.New("failed to get request")).Error(),
		})
	}

	statusCode, err := r.subDistrictService.ImportSubDistricts(req)
	if err != nil {
		return c.Status(statusCode).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	return c.Status(statusCode).JSON(pkg.SuccessResponse{
		Message: "sub-districts imported successfully",
		Data:    nil,
	})
}

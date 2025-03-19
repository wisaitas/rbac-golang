package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/requests"
	provinceService "github.com/wisaitas/rbac-golang/internal/auth-service/services/province"
	"github.com/wisaitas/rbac-golang/pkg"
)

type ProvinceHandler struct {
	provinceService provinceService.ProvinceService
}

func NewProvinceHandler(
	provinceService provinceService.ProvinceService,
) *ProvinceHandler {
	return &ProvinceHandler{
		provinceService: provinceService,
	}
}

func (r *ProvinceHandler) GetProvinces(c *fiber.Ctx) error {
	query, ok := c.Locals("query").(pkg.PaginationQuery)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(errors.New("failed to get queries")).Error(),
		})
	}

	provinces, statusCode, err := r.provinceService.GetProvinces(query)
	if err != nil {
		return c.Status(statusCode).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(pkg.SuccessResponse{
		Message: "provinces fetched successfully",
		Data:    provinces,
	})
}

func (r *ProvinceHandler) ImportProvinces(c *fiber.Ctx) error {
	req, ok := c.Locals("request").(requests.ImportProvince)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(errors.New("failed to get request")).Error(),
		})
	}

	statusCode, err := r.provinceService.ImportProvinces(req)
	if err != nil {
		return c.Status(statusCode).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	return c.Status(statusCode).JSON(pkg.SuccessResponse{
		Message: "provinces imported successfully",
		Data:    req,
	})
}

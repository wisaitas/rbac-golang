package validates

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/queries"
	"github.com/wisaitas/rbac-golang/pkg"
)

type DistrictValidate interface {
	ValidateGetDistrictsRequest(c *fiber.Ctx) error
}

type districtValidate struct {
	validatorUtil pkg.ValidatorUtil
}

func NewDistrictValidate(validatorUtil pkg.ValidatorUtil) DistrictValidate {
	return &districtValidate{
		validatorUtil: validatorUtil,
	}
}

func (r *districtValidate) ValidateGetDistrictsRequest(c *fiber.Ctx) error {
	query := queries.DistrictQuery{}

	if err := validateCommonPaginationQuery(c, &query, r.validatorUtil); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	c.Locals("query", query)
	return c.Next()
}

package validates

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/queries"
	"github.com/wisaitas/rbac-golang/pkg"
)

type SubDistrictValidate interface {
	ValidateGetSubDistrictsRequest(c *fiber.Ctx) error
}

type subDistrictValidate struct {
	validatorUtil pkg.ValidatorUtil
}

func NewSubDistrictValidate(validatorUtil pkg.ValidatorUtil) SubDistrictValidate {
	return &subDistrictValidate{
		validatorUtil: validatorUtil,
	}
}

func (r *subDistrictValidate) ValidateGetSubDistrictsRequest(c *fiber.Ctx) error {
	query := queries.SubDistrictQuery{}

	if err := validateCommonPaginationQuery(c, &query, r.validatorUtil); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	c.Locals("query", query)
	return c.Next()
}

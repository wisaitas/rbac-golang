package validates

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/queries"
	"github.com/wisaitas/rbac-golang/pkg"
)

type SubDistrictValidate struct {
	validatorUtil pkg.ValidatorUtil
}

func NewSubDistrictValidate(validatorUtil pkg.ValidatorUtil) *SubDistrictValidate {
	return &SubDistrictValidate{
		validatorUtil: validatorUtil,
	}
}

func (r *SubDistrictValidate) ValidateGetSubDistrictsRequest(c *fiber.Ctx) error {
	query := queries.SubDistrictQuery{}

	if err := validateCommonPaginationQuery(c, &query, r.validatorUtil); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	c.Locals("query", query)
	return c.Next()
}

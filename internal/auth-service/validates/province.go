package validates

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/rbac-golang/pkg"
)

type ProvinceValidate interface {
	ValidateGetProvincesRequest(c *fiber.Ctx) error
}

type provinceValidate struct {
	validatorUtil pkg.ValidatorUtil
}

func NewProvinceValidate(validatorUtil pkg.ValidatorUtil) ProvinceValidate {
	return &provinceValidate{
		validatorUtil: validatorUtil,
	}
}

func (r *provinceValidate) ValidateGetProvincesRequest(c *fiber.Ctx) error {
	query := pkg.PaginationQuery{}

	if err := validateCommonPaginationQuery(c, &query, r.validatorUtil); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	c.Locals("query", query)
	return c.Next()
}

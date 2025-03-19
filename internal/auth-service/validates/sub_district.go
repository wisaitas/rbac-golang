package validates

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/queries"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/requests"
	"github.com/wisaitas/rbac-golang/pkg"
)

type SubDistrictValidate struct {
}

func NewSubDistrictValidate() *SubDistrictValidate {
	return &SubDistrictValidate{}
}

func (r *SubDistrictValidate) ValidateGetSubDistrictsRequest(c *fiber.Ctx) error {
	query := queries.SubDistrictQuery{}

	if err := validateCommonPaginationQuery(c, &query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	c.Locals("query", query)
	return c.Next()
}

func (r *SubDistrictValidate) ValidateImportSubDistrictsRequest(c *fiber.Ctx) error {
	request := requests.ImportSubDistrict{}

	if err := validateCommonRequestFormBody(c, &request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	c.Locals("req", request)
	return c.Next()
}

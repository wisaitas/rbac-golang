package validates

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/requests"
	"github.com/wisaitas/rbac-golang/pkg"
)

type DistrictValidate struct {
}

func NewDistrictValidate() *DistrictValidate {
	return &DistrictValidate{}
}

func (r *DistrictValidate) ValidateGetDistrictsRequest(c *fiber.Ctx) error {
	query := pkg.PaginationQuery{}

	if err := validateCommonPaginationQuery(c, &query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	c.Locals("query", query)
	return c.Next()
}

func (r *DistrictValidate) ValidateImportDistrictsRequest(c *fiber.Ctx) error {
	request := requests.ImportDistrict{}

	if err := validateCommonRequestFormBody(c, &request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	c.Locals("req", request)
	return c.Next()
}

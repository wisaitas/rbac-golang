package validates

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/requests"
	"github.com/wisaitas/rbac-golang/pkg"
)

type ProvinceValidate struct {
}

func NewProvinceValidate() *ProvinceValidate {
	return &ProvinceValidate{}
}

func (r *ProvinceValidate) ValidateGetProvincesRequest(c *fiber.Ctx) error {
	query := pkg.PaginationQuery{}

	if err := validateCommonPaginationQuery(c, &query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	c.Locals("query", query)
	return c.Next()
}

func (r *ProvinceValidate) ValidateImportProvincesRequest(c *fiber.Ctx) error {
	request := requests.ImportProvince{}

	if err := validateCommonRequestFormBody(c, &request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	c.Locals("request", request)
	return c.Next()
}

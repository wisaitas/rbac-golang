package validates

import (
	"fmt"

	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/params"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/requests"
	"github.com/wisaitas/rbac-golang/pkg"

	"github.com/gofiber/fiber/v2"
)

type UserValidate interface {
	ValidateCreateUserRequest(c *fiber.Ctx) error
	ValidateGetUsersRequest(c *fiber.Ctx) error
	ValidateUpdateUserRequest(c *fiber.Ctx) error
}

type userValidate struct {
	validatorUtil pkg.ValidatorUtil
}

func NewUserValidate(validatorUtil pkg.ValidatorUtil) UserValidate {
	return &userValidate{
		validatorUtil: validatorUtil,
	}
}

func (r *userValidate) ValidateCreateUserRequest(c *fiber.Ctx) error {
	req := requests.CreateUserRequest{}

	if err := validateCommonRequestJSONBody(c, &req, r.validatorUtil); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: fmt.Sprintf("failed to validate request: %s", err.Error()),
		})
	}

	c.Locals("req", req)
	return c.Next()
}

func (r *userValidate) ValidateGetUsersRequest(c *fiber.Ctx) error {
	query := pkg.PaginationQuery{}

	if err := validateCommonPaginationQuery(c, &query, r.validatorUtil); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	c.Locals("query", query)
	return c.Next()

}

func (r *userValidate) ValidateUpdateUserRequest(c *fiber.Ctx) error {
	req := requests.UpdateUserRequest{}
	params := params.UserParams{}

	if err := validateCommonRequestParams(c, &params, r.validatorUtil); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	if err := validateCommonRequestJSONBody(c, &req, r.validatorUtil); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	c.Locals("req", req)
	c.Locals("params", params)
	return c.Next()
}

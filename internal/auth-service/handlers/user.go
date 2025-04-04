package handlers

import (
	"errors"

	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/params"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/requests"
	"github.com/wisaitas/rbac-golang/internal/auth-service/models"
	"github.com/wisaitas/rbac-golang/internal/auth-service/services/user"
	"github.com/wisaitas/rbac-golang/pkg"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService user.UserService
}

func NewUserHandler(
	userService user.UserService,
) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (r *UserHandler) GetUsers(c *fiber.Ctx) error {
	query, ok := c.Locals("query").(pkg.PaginationQuery)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(errors.New("failed to get queries")).Error(),
		})
	}

	users, statusCode, err := r.userService.GetUsers(query)
	if err != nil {
		return c.Status(statusCode).JSON(pkg.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(statusCode).JSON(pkg.SuccessResponse{
		Message: "users fetched successfully",
		Data:    users,
	})
}

func (r *UserHandler) GetUserProfile(c *fiber.Ctx) error {
	userContext, ok := c.Locals("userContext").(models.UserContext)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(errors.New("failed to get user context")).Error(),
		})
	}

	resp, statusCode, err := r.userService.GetUserProfile(userContext)
	if err != nil {
		return c.Status(statusCode).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	return c.Status(statusCode).JSON(pkg.SuccessResponse{
		Message: "user profile fetched successfully",
		Data:    resp,
	})
}

func (r *UserHandler) CreateUser(c *fiber.Ctx) error {
	req, ok := c.Locals("req").(requests.CreateUserRequest)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(errors.New("failed to get request")).Error(),
		})
	}

	user, statusCode, err := r.userService.CreateUser(req)
	if err != nil {
		return c.Status(statusCode).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	return c.Status(statusCode).JSON(pkg.SuccessResponse{
		Message: "user created successfully",
		Data:    user,
	})
}

func (r *UserHandler) UpdateUser(c *fiber.Ctx) error {
	req, ok := c.Locals("req").(requests.UpdateUserRequest)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(errors.New("failed to get request")).Error(),
		})
	}

	param, ok := c.Locals("params").(params.UserParams)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(errors.New("failed to get params")).Error(),
		})
	}

	return c.JSON(c.Locals("userContext")) // test bug for sonarqube scan found

	resp, statusCode, err := r.userService.UpdateUser(param, req)
	if err != nil {
		return c.Status(statusCode).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	return c.Status(statusCode).JSON(pkg.SuccessResponse{
		Message: "user updated successfully",
		Data:    resp,
	})
}

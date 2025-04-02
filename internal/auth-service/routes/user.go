package routes

import (
	"github.com/wisaitas/rbac-golang/internal/auth-service/handlers"
	"github.com/wisaitas/rbac-golang/internal/auth-service/middlewares"
	"github.com/wisaitas/rbac-golang/internal/auth-service/validates"

	"github.com/gofiber/fiber/v2"
)

type UserRoutes struct {
	app            fiber.Router
	userHandler    *handlers.UserHandler
	userValidate   *validates.UserValidate
	authMiddleware *middlewares.AuthMiddleware
	userMiddleware *middlewares.UserMiddleware
}

func NewUserRoutes(
	app fiber.Router,
	userHandler *handlers.UserHandler,
	userValidate *validates.UserValidate,
	authMiddleware *middlewares.AuthMiddleware,
	userMiddleware *middlewares.UserMiddleware,
) *UserRoutes {
	return &UserRoutes{
		app:            app,
		userHandler:    userHandler,
		userValidate:   userValidate,
		authMiddleware: authMiddleware,
		userMiddleware: userMiddleware,
	}
}

func (r *UserRoutes) UserRoutes() {
	userRoutes := r.app.Group("/users")

	// Method GET
	userRoutes.Get("/", r.userValidate.ValidateGetUsersRequest, r.userHandler.GetUsers)
	userRoutes.Get("/profile", r.userMiddleware.GetUserProfile, r.userHandler.GetUserProfile)

	// Method POST
	userRoutes.Post("/", r.userValidate.ValidateCreateUserRequest, r.userHandler.CreateUser)

	// Method PATCH
	userRoutes.Patch("/:id", r.userMiddleware.UpdateUser, r.userValidate.ValidateUpdateUserRequest, r.userHandler.UpdateUser)
}

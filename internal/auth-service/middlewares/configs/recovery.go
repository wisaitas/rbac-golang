package configs

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/wisaitas/rbac-golang/pkg"
)

func Recovery() fiber.Handler {
	return recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			log.Println(e)
			c.Status(fiber.StatusInternalServerError).JSON(pkg.ErrorResponse{
				Message: pkg.Error(errors.New("internal server error")).Error(),
			})
		},
	})
}

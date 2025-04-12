package initial

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/wisaitas/rbac-golang/internal/auth-service/env"
	"github.com/wisaitas/rbac-golang/pkg"

	"github.com/gofiber/fiber/v2"
)

func init() {
	env.LoadEnv()
}

func InitializeApp() {
	app := fiber.New()

	setupMiddlewares(app)

	config := initializeConfig()

	util := initializeUtil(config)

	repo := initializeRepositorie(config)
	services := initializeService(repo, util)
	handlers := initializeHandler(services)
	validate := initializeValidate(util)
	middleware := initializeMiddleware(util)

	initializeRoute(app, handlers, validate, middleware)

	run(app, config)
}

func run(app *fiber.App, config *config) {
	go func() {
		if err := app.Listen(fmt.Sprintf(":%s", env.ENV.PORT)); err != nil {
			log.Fatalf("error starting server: %v\n", pkg.Error(err))
		}
	}()

	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-gracefulShutdown

	close(app, config)
}

func close(app *fiber.App, config *config) {
	sqlDB, err := config.DB.DB()
	if err != nil {
		log.Fatalf("error getting database: %v\n", pkg.Error(err))
	}

	if err := sqlDB.Close(); err != nil {
		log.Fatalf("error closing database: %v\n", pkg.Error(err))
	}

	if err := config.Redis.Close(); err != nil {
		log.Fatalf("error closing redis: %v\n", pkg.Error(err))
	}

	if err := app.Shutdown(); err != nil {
		log.Fatalf("error shutting down app: %v\n", pkg.Error(err))
	}

	log.Println("gracefully shutdown")
}

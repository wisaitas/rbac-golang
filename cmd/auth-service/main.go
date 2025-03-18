package main

import "github.com/wisaitas/rbac-golang/internal/auth-service/initial"

func main() {
	app := initial.InitializeApp()

	app.SetupMiddlewares()

	app.SetupRoutes()

	app.Run()
}

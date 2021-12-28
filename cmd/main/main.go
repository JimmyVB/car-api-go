package main

import (
	"car-api/api"
	"car-api/internal"
	"github.com/gofiber/fiber"
)

func main() {

	app := fiber.New()

	internal.SetErrorHandler(app)
	api.SetupPokemonsRoutes(app)
	api.SetupUsersRouters(app)

	_ = app.Listen("3000")
}

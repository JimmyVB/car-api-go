package main

import (
	"car-api/api"
	"car-api/internal"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

func main() {

	app := fiber.New()
	key := "tokenKey"

	internal.SetErrorHandler(app)
	app.Use(middleware.Recover())
	api.SetupPokemonsRoutes(app, key)
	api.SetupUsersRouters(app, key)

	_ = app.Listen("3000")
}

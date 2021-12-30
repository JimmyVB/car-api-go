package routes

import (
	"car-api/internal/middleware"
	"car-api/internal/service"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

// UsersRouters func for describe token routes.
func UsersRouters(app *fiber.App, tokenKey string) {

	//Inyectando el servicio
	s := service.Start(tokenKey)

	// Create routes group.
	route := app.Group("/api/v1")

	// Routes
	route.Post("/create", s.CreateUserHandler)
	route.Post("/login", s.LoginHandler)
}

// CarRouters func for describe cars routes.
func CarRouters(app *fiber.App, tokenKey string) {
	s := service.Start(tokenKey)
	route := app.Group("/api/v1/cars")
	route.Use(middleware.JwtMiddleware(tokenKey)).Post("/create", s.CreateHandler)
	route.Get("/all", s.GetAllHandler)
	route.Get("/find/:id", s.GetOneHandler)
	route.Put("/update/:id", s.UpdateHandler)
	//route.Delete("/delete", s.DeleteHandler)
	route.Use(middleware.JwtMiddleware(tokenKey)).Get("/fin", s.CreateHandler)

}

// SwaggerRouters func for describe group of API Docs routes.
func SwaggerRouters(app *fiber.App) {
	route := app.Group("/swagger")
	route.Get("*", swagger.Handler)
}

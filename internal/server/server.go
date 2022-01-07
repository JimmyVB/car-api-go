package server

import (
	"car-api/internal/core/ports"
	"car-api/internal/middleware"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"log"
)

type Server struct {
	userHandler ports.IUserHandler
	carHandler  ports.ICarHandler
}

func NewServer(uHandler ports.IUserHandler, cHandler ports.ICarHandler) *Server {
	return &Server{
		userHandler: uHandler,
		carHandler:  cHandler,
	}
}

func (s *Server) Initialize() {
	app := fiber.New()

	app.Use(recover2.New())

	routes := app.Group("/swagger")
	routes.Get("*", swagger.Handler)

	routes = app.Group("/api/v1")
	routes.Post("/login", s.userHandler.Login)
	routes.Post("/register", s.userHandler.SaveUser)

	routes = app.Group("/api/v1/cars")
	routes.Get("/all", s.carHandler.GetAll)
	routes.Get("/find/:id", s.carHandler.GetOne)

	routes.Use(middleware.JwtMiddleware("secret")).Post("/create", s.carHandler.Save)
	routes.Put("/update/:id", s.carHandler.Update)
	routes.Delete("/delete/:id", s.carHandler.Delete)

	err := app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

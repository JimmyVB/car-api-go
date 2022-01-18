package server

import (
	enums "car-api/internal/core/emuns"
	"car-api/internal/core/ports"
	"car-api/internal/middleware"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"os"
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
	app.Use(cors.New())
	// Set config for logger
	file := s.LoggingMiddleware()
	defer file.Close()
	loggerConfig := logger.Config{
		Output: file, // add file to save output
	}
	app.Use(logger.New(loggerConfig))
	routes := app.Group(enums.RouterSwagger)
	routes.Get(enums.RouterAny, swagger.Handler)

	routes = app.Group(enums.RouterGroupGlobal)
	routes.Post(enums.RouterLogin, s.userHandler.Login)
	routes.Post(enums.RouterRegister, s.userHandler.SaveUser)

	routes = app.Group(enums.RouterCarsGroup)
	routes.Get("/", s.carHandler.GetAll)
	routes.Get(enums.RouterParamID, s.carHandler.GetOne)

	routes.Use(middleware.JwtMiddleware("secret")).Post("/", s.carHandler.Save)
	routes.Put(enums.RouterParamID, s.carHandler.Update)
	routes.Delete(enums.RouterParamID, s.carHandler.Delete)

	port := os.Getenv("PORT")

	err := app.Listen(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Server) LoggingMiddleware() *os.File {
	// Define file to logs
	file, err := os.OpenFile("./logs/logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	return file
}

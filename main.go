package main

import (
	_ "car-api/docs"
	"car-api/internal/core/services"
	"car-api/internal/handler"
	"car-api/internal/repository"
	"car-api/internal/server"
	"database/sql"
	"fmt"
	_ "github.com/arsmn/fiber-swagger/v2"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"log"
	"time"
)

// @title Car API
// @version 1.0
// @description This is a small CRUD in Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name Jimmy Valdez
// @contact.email jimmyvb16@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {

	var cfg config
	err := envconfig.Process("CAR", &cfg)
	postgresUri := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName)

	db, err := sql.Open("postgres", postgresUri)

	if err != nil {
		log.Fatal(err)
	}

	userRepository := repository.NewUserRepository(db)
	carRepository := repository.NewCarRepository(db)
	//services
	userService := services.NewUserService(userRepository)
	carService := services.NewCarService(carRepository)
	//handlers
	userHandler := handler.NewUserHandler(userService)
	carHandler := handler.NewCarHandler(carService)
	//server
	httpServer := server.NewServer(
		userHandler,
		carHandler,
	)
	httpServer.Initialize()
}

type config struct {

	// Database configuration
	DbUser    string        `default:"xzmbprjsejsazz"`
	DbPass    string        `default:"1b2e39d2b1b6d7098cf2a756a4706276817729749cde329086b2772ee1f9d74a"`
	DbHost    string        `default:"ec2-3-224-157-224.compute-1.amazonaws.com"`
	DbPort    string        `default:"5432"`
	DbName    string        `default:"dbiqfdis1j5q6g"`
	DbTimeout time.Duration `default:"5s"`
}

package main

import (
	"car-api/internal/core/domain"
	"car-api/internal/core/services"
	"car-api/internal/repository"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kelseyhightower/envconfig"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
	"strconv"
	"time"
)

// Servicio / Actividad / Lógica de negocio
func create(ctx context.Context, car domain.Car) (*domain.Car, error) {
	carService := Initialize()
	logger := activity.GetLogger(ctx)
	logger.Info("Crear vehiculo : " + car.Mark + "-" + car.Model) // para este caso consideramos el print como transacción exitosa
	err := carService.Save(car)
	if err != nil {
		return nil, err
	}

	return &car, nil // auto recien creado
}

// Servicio / Actividad / Lógica de negocio
func getAll(ctx context.Context, carService services.CarService) ([]domain.Car, error) {

	logger := activity.GetLogger(ctx)
	logger.Info("Obtener listado de vehiculos") // para este caso consideramos el print como transacción exitosa
	cars, err := carService.GetAll()

	if err != nil {
		return nil, err
	}

	return cars, nil // listado de vehiculos
}

// Servicio / Actividad / Lógica de negocio
func getOne(ctx context.Context, carService services.CarService, id string) (*domain.Car, error) {

	logger := activity.GetLogger(ctx)
	logger.Info("Obtener vehiculo: " + id) // para este caso consideramos el print como transacción exitosa

	car, err := carService.GetOne(id)

	if err != nil {
		return nil, err
	}

	return &car, nil // vehiculo
}

// Servicio / Actividad / Lógica de negocio
func update(ctx context.Context, carService services.CarService, id string, car domain.Car) (*domain.Car, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Update vehiculo : " + car.Mark + "-" + car.Model) // para este caso consideramos el print como transacción exitosa
	err := carService.Update(id, car)
	if err != nil {
		return nil, err
	}

	return &car, nil // auto actualizado
}

// Servicio / Actividad / Lógica de negocio
func delete(ctx context.Context, carService services.CarService, id string) error {

	logger := activity.GetLogger(ctx)
	logger.Info("Delete vehiculo: " + id) // para este caso consideramos el print como transacción exitosa

	err := carService.Delete(id)

	return err // vehiculo
}

// Servicio / Actividad / Lógica de negocio
func rent(ctx context.Context, carRent *domain.CarRent) (*domain.CarRent, error) {
	carService := Initialize()
	logger := activity.GetLogger(ctx)
	logger.Info("Rentar vehiculo : " + strconv.Itoa(carRent.IdCar)) // para este caso consideramos el print como transacción exitosa
	res, err := carService.RentCar(carRent)
	if err != nil {
		return nil, err
	}

	return res, nil // auto rentado
}

func main() {

	// Para iniciar debes abrir un objeto client. Este es un objeto pesado
	// Solo debes crear uno por proceso.
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("No es posible crear el cliente", err)
	}
	defer c.Close()

	// creo el worker

	// como es el microservicio de vehiculo, tendrá un taskqueue llamado vehiculo

	// El taskqueue es importante porque nos permite enrutar las tareas a los workers

	w := worker.New(c, "car", worker.Options{}) // TaskQueue= car

	// registramos el identificador para la acción de create de vehiculo
	oa1 := activity.RegisterOptions{
		Name: "car.create", // identificador del servicio / actividad. Puede ser cualquier string que sea único para la aplicación
	}

	// asociamos el identificador con la función create
	w.RegisterActivityWithOptions(create, oa1) // registro el servicio / actividad con el worker

	// registramos el identificador para la acción de obtener listado vehiculos
	oa2 := activity.RegisterOptions{
		Name: "car.getAll", // identificador del servicio / actividad. Puede ser cualquier string que sea único para la aplicación
	}

	// asociamos el identificador con la función getAll
	w.RegisterActivityWithOptions(getAll, oa2) // registro el servicio / actividad con el worker

	// registramos el identificador para la acción de update vehiculo
	oa3 := activity.RegisterOptions{
		Name: "car.update", // identificador del servicio / actividad. Puede ser cualquier string que sea único para la aplicación
	}

	// asociamos el identificador con la función update
	w.RegisterActivityWithOptions(update, oa3) // registro el servicio / actividad con el worker

	// registramos el identificador para la acción de obtener vehiculo
	oa4 := activity.RegisterOptions{
		Name: "car.getOne", // identificador del servicio / actividad. Puede ser cualquier string que sea único para la aplicación
	}

	// asociamos el identificador con la función getOne
	w.RegisterActivityWithOptions(getOne, oa4) // registro el servicio / actividad con el worker

	// registramos el identificador para la acción de delete vehiculo
	oa5 := activity.RegisterOptions{
		Name: "car.delete", // identificador del servicio / actividad. Puede ser cualquier string que sea único para la aplicación
	}

	// asociamos el identificador con la función delete
	w.RegisterActivityWithOptions(delete, oa5) // registro el servicio / actividad con el worker

	// registramos el identificador para la acción de delete vehiculo
	oa6 := activity.RegisterOptions{
		Name: "car.rent", // identificador del servicio / actividad. Puede ser cualquier string que sea único para la aplicación
	}

	// asociamos el identificador con la función rent
	w.RegisterActivityWithOptions(rent, oa6) // registro el servicio / actividad con el worker

	// Ejecutar worker.
	// Este es un proceso demonio.
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("No es posible ejecutar worker", err)
	}
}

func Initialize() *services.CarService {
	//Inicializar service de car
	var cfg config
	err := envconfig.Process("CAR", &cfg)
	postgresUri := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=require",
		cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName)

	db, err := sql.Open("postgres", postgresUri)

	if err != nil {
		log.Fatal(err)
	}

	//repository
	carRepository := repository.NewCarRepository(db)
	//service
	carService := services.NewCarService(carRepository)

	return carService
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

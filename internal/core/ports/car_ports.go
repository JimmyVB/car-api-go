package ports

import (
	user "car-api/internal/core/domain"
	"github.com/gofiber/fiber/v2"
)

//go:generate mockery --case=snake --outpkg=mocks --output=mocks --name=ICarRepository

type ICarService interface {
	Save(car user.Car) error
	GetAll() ([]user.Car, error)
	GetOne(id string) (user.Car, error)
	Update(id string, car user.Car) error
	Delete(id string) error
	RentCar(carRent *user.CarRent) (*user.CarRent, error)
}

type ICarRepository interface {
	Save(car user.Car) error
	GetAll() ([]user.Car, error)
	GetOne(id string) (user.Car, error)
	Update(id string, car user.Car) error
	Delete(id string) error
	RentCar(carRent *user.CarRent) (*user.CarRent, error)
}

type ICarHandler interface {
	Save(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetOne(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

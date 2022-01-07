package services

import (
	user "car-api/internal/core/domain"
	"car-api/internal/core/ports"
	"github.com/gofiber/fiber/v2"
)

type CarService struct {
	carRepostiory ports.ICarRepository
}

func NewCarService(repository ports.ICarRepository) *CarService {
	return &CarService{
		carRepostiory: repository,
	}
}

func (c *CarService) Save(car user.Car) error {
	err := c.carRepostiory.Save(car)
	if err != nil {
		return fiber.NewError(400, "cannot create car")
	}
	return nil
}

func (c *CarService) GetAll() ([]user.Car, error) {
	res, err := c.carRepostiory.GetAll()
	if err != nil {
		return nil, fiber.NewError(400, "cannot get all car")
	}
	return res, nil
}

func (c *CarService) GetOne(id string) (user.Car, error) {
	res, err := c.carRepostiory.GetOne(id)
	if err != nil {
		return user.Car{}, fiber.NewError(400, "cannot get one car")
	}
	return res, nil
}

func (c *CarService) Update(id string, car user.Car) error {
	err := c.carRepostiory.Update(id, car)
	if err != nil {
		return fiber.NewError(400, "cannot update car")
	}
	return nil
}

func (c *CarService) Delete(id string) error {
	err := c.carRepostiory.Delete(id)
	if err != nil {
		return fiber.NewError(400, "cannot delete car")
	}
	return nil
}

var _ ports.ICarService = (*CarService)(nil)

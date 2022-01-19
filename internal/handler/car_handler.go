package handler

import (
	"car-api/internal/core/domain"
	enums "car-api/internal/core/emuns"
	"car-api/internal/core/ports"
	"car-api/internal/kafka"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type CarHandler struct {
	carService ports.ICarService
}

// Save method for create a new car.
// @Description create a new car.
// @Summary create a new car
// @Tags Car
// @Accept json
// @Produce json
// @Param request body domain.Car true "Create Car Data"
// @Success 200 {object} domain.Message
// @Success 400 {object} domain.MessageError
// @Success 500 {object} domain.MessageError
// @Security ApiKeyAuth
// @Router /cars [post]
func (ch *CarHandler) Save(c *fiber.Ctx) error {
	var car domain.Car
	err := c.BodyParser(&car)
	if err != nil {
		return fiber.NewError(400, "cannot parse params")
	}

	err = ch.carService.Save(car)
	if err != nil {
		return fiber.NewError(404, "cannot create a car")
	}
	carEvent := kafka.CarEvent{Operation: "create"}
	eventInBytes, err := json.Marshal(carEvent)
	kafka.PushMessageToQueue(eventInBytes)
	return c.JSON(domain.Message{
		Message: fmt.Sprintf("%s %s %s", enums.Created, enums.Successfully, enums.Car),
	})
}

// GetAll method for get all car.
// @Description get all car.
// @Summary get all car
// @Tags Car
// @Accept json
// @Produce json
// @Success 200 {object} domain.Message
// @Success 400 {object} domain.MessageError
// @Success 500 {object} domain.MessageError
// @Router /cars [get]
func (ch *CarHandler) GetAll(c *fiber.Ctx) error {

	cars, err := ch.carService.GetAll()

	if err != nil {
		return fiber.NewError(400, "cannot get all car")
	}

	return c.JSON(domain.Message{
		Message: fmt.Sprintf("%s %s %s", enums.Loaded, enums.Successfully, enums.Cars),
		Data:    cars,
	})

}

// GetOne method for get one car.
// @Description get one car.
// @Summary get one car
// @Tags Car
// @Accept json
// @Produce json
// @Param id path string true "Car ID"
// @Success 200 {object} domain.Message
// @Success 400 {object} domain.MessageError
// @Success 500 {object} domain.MessageError
// @Router /cars/{id} [get]
func (ch *CarHandler) GetOne(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return fiber.NewError(400, "cannot parse ID")
	}

	car, err := ch.carService.GetOne(id)

	if err != nil {
		return fiber.NewError(400, "cannot get one car")
	}

	return c.JSON(domain.Message{
		Message: fmt.Sprintf("%s %s %s", enums.Loaded, enums.Successfully, enums.Car),
		Data:    car,
	})
}

// Update method for update one car.
// @Description update one car.
// @Summary update one car
// @Tags Car
// @Accept json
// @Produce json
// @Param id path string true "Car ID"
// @Param request body domain.Car true "Update Car Data"
// @Success 200 {object} domain.Message
// @Success 400 {object} domain.MessageError
// @Success 500 {object} domain.MessageError
// @Security ApiKeyAuth
// @Router /cars/{id} [put]
func (ch *CarHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return fiber.NewError(400, "cannot parse ID")
	}

	var car domain.Car
	err := c.BodyParser(&car)

	if err != nil {
		return fiber.NewError(400, "cannot body parser")
	}

	err = ch.carService.Update(id, car)

	if err != nil {
		return fiber.NewError(400, "cannot update a car")
	}
	carEvent := kafka.CarEvent{Operation: "update"}
	eventInBytes, err := json.Marshal(carEvent)
	kafka.PushMessageToQueue(eventInBytes)
	return c.JSON(domain.Message{
		Message: fmt.Sprintf("%s %s %s", enums.Updated, enums.Successfully, enums.Car),
	})
}

// Delete method for delete one car.
// @Description delete one car.
// @Summary delete one car
// @Tags Car
// @Accept json
// @Produce json
// @Param id path string true "Car ID"
// @Success 200 {object} domain.Message
// @Success 400 {object} domain.MessageError
// @Success 500 {object} domain.MessageError
// @Security ApiKeyAuth
// @Router /cars/{id} [delete]
func (ch *CarHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return fiber.NewError(400, "cannot parse ID")
	}

	err := ch.carService.Delete(id)

	if err != nil {
		return fiber.NewError(400, "cannot delete car")
	}
	carEvent := kafka.CarEvent{Operation: "delete"}
	eventInBytes, err := json.Marshal(carEvent)
	kafka.PushMessageToQueue(eventInBytes)
	return c.JSON(domain.Message{
		Message: fmt.Sprintf("%s %s %s", enums.Deleted, enums.Successfully, enums.Car),
	})
}

var _ ports.ICarHandler = (*CarHandler)(nil)

func NewCarHandler(carService ports.ICarService) *CarHandler {
	return &CarHandler{
		carService: carService,
	}
}

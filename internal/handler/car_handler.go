package handler

import (
	user "car-api/internal/core/domain"
	"car-api/internal/core/ports"
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
// @Success 200 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/cars/create [post]
func (ch *CarHandler) Save(c *fiber.Ctx) error {
	var car user.Car
	err := c.BodyParser(&car)
	if err != nil {
		return fiber.NewError(400, "cannot parse params")
	}

	err = ch.carService.Save(car)
	if err != nil {
		return fiber.NewError(404, "cannot create a car")
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "car added",
	})
}

// GetAll method for get all car.
// @Description get all car.
// @Summary get all car
// @Tags Car
// @Accept json
// @Produce json
// @Success 200 {array} user.Car
// @Router /v1/cars/all [get]
func (ch *CarHandler) GetAll(c *fiber.Ctx) error {

	cars, err := ch.carService.GetAll()

	if err != nil {
		return fiber.NewError(400, "cannot get all car")
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "cars obtained",
		"cars":  cars,
	})
}

// GetOne method for get one car.
// @Description get one car.
// @Summary get one car
// @Tags Car
// @Accept json
// @Produce json
// @Param id path string true "Car ID"
// @Success 200 {object} user.Car
// @Router /v1/cars/find/{id} [get]
func (ch *CarHandler) GetOne(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return fiber.NewError(400, "cannot parse ID")
	}

	car, err := ch.carService.GetOne(id)

	if err != nil {
		return fiber.NewError(400, "cannot get one car")
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "Get one car",
		"car":   car,
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
// @Success 200 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/cars/update/{id} [put]
func (ch *CarHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return fiber.NewError(400, "cannot parse ID")
	}

	var car user.Car
	err := c.BodyParser(&car)

	if err != nil {
		return fiber.NewError(400, "cannot body parser")
	}

	err = ch.carService.Update(id, car)

	if err != nil {
		return fiber.NewError(400, "cannot update a car")
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "A car was updated",
	})
}

// Delete method for delete one car.
// @Description delete one car.
// @Summary delete one car
// @Tags Car
// @Accept json
// @Produce json
// @Param id path string true "Car ID"
// @Success 200 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/cars/delete/{id} [delete]
func (ch *CarHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return fiber.NewError(400, "cannot parse ID")
	}

	err := ch.carService.Delete(id)

	if err != nil {
		return fiber.NewError(400, "cannot delete car")
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "Delete one car",
	})
}

var _ ports.ICarHandler = (*CarHandler)(nil)

func NewCarHandler(carService ports.ICarService) *CarHandler {
	return &CarHandler{
		carService: carService,
	}
}

package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
	"github.com/google/uuid"
)

// CreateHandler method for create a new car.
// @Description create a new car.
// @Summary create a new car
// @Tags Car
// @Accept json
// @Produce json
// @Param marca body string true "Marca"
// @Param model body string true "Modelo"
// @Param price body string true "Precio"
// @Success 200 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/cars/create [post]
func (w *WebServices) CreateHandler(c *fiber.Ctx) error {

	var cmd carCMD
	_ = c.BodyParser(&cmd)

	err := w.cars.Save(utils.UUID(), cmd.Mark, cmd.Model, cmd.Price)

	if err != nil {
		return fiber.NewError(400, "cannot add to car")
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "car added",
	})
}

// GetAllHandler method for get all car.
// @Description get all car.
// @Summary get all car
// @Tags Car
// @Accept json
// @Produce json
// @Success 200 {array} carCMD
// @Router /v1/cars/all [get]
func (w *WebServices) GetAllHandler(c *fiber.Ctx) error {

	var cmd carCMD
	_ = c.BodyParser(&cmd)

	cars, err := w.cars.GetAll()

	if err != nil {
		return fiber.NewError(400, "cannot get all car")
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "car added",
		"cars":  cars,
	})
}

// GetOneHandler method for get one car.
// @Description get one car.
// @Summary get one car
// @Tags Car
// @Accept json
// @Produce json
// @Param id path string true "Car ID"
// @Success 200 {object} carCMD
// @Router /v1/cars/find/{id} [get]
func (w *WebServices) GetOneHandler(c *fiber.Ctx) error {

	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return fiber.NewError(400, "cannot get one car")
	}

	cars, err := w.cars.GetOne(id)

	if err != nil {
		return fiber.NewError(400, "cannot get all car")
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "Get one car",
		"car":   cars,
	})
}

// UpdateHandler method for update one car.
// @Description update one car.
// @Summary update one car
// @Tags Car
// @Accept json
// @Produce json
// @Param marca body string true "Marca"
// @Param model body string true "Modelo"
// @Param price body integer true "Precio"
// @Success 200 {object} carCMD
// @Security ApiKeyAuth
// @Router /v1/cars/update/{id} [put]
func (w *WebServices) UpdateHandler(c *fiber.Ctx) error {

	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return fiber.NewError(400, "cannot ID")
	}

	var cmd CreateCarCMD
	err = c.BodyParser(&cmd)

	if err != nil {
		return fiber.NewError(400, "cannot body parser")
	}

	car, err := w.cars.Update(id, cmd)

	if err != nil {
		return fiber.NewError(400, "cannot get all car")
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "Get one car",
		"car":   car,
	})
}

type carCMD struct {
	ID    string `json:"id"`
	Mark  string `json:"marca"`
	Model string `json:"model"`
	Price uint   `json:"price"`
}

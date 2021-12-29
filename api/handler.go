package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
)

func (w *WebServices) CreateHandler(c *fiber.Ctx) error {

	var cmd carCMD
	_ = c.BodyParser(&cmd)

	err := w.cars.Save(utils.UUID(), cmd.Mark, cmd.Model, cmd.Price)

	if err != nil {
		return fiber.NewError(400, "cannot add to car")
	}

	return c.JSON(struct {
		Res string `json:"result"`
	}{
		Res: "car added",
	})
}

type carCMD struct {
	ID    string `json:"id"`
	Mark  string `json:"mark"`
	Model string `json:"model"`
	Price uint   `json:"price"`
}

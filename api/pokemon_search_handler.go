package api

import (
	"github.com/gofiber/fiber/v2"
)

func (w *WebServices) SearchPokemonHandler(c *fiber.Ctx) error {

	res, err := w.search.Search(PokemonFilter{
		Name: c.Query("name"),
		Type: c.Query("type"),
	})

	if err != nil {
		return fiber.NewError(400, "cannot bring pokemons")
	}

	if len(res) == 0 {
		_ = c.JSON([]interface{}{})
	}

	return c.JSON(res)
}

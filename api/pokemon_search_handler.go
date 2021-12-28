package api

import (
	"github.com/gofiber/fiber"
)

func (w *WebServices) SearchPokemonHandler(c *fiber.Ctx) {

	res, err := w.search.Search(PokemonFilter{
		Name: c.Query("name"),
		Type: c.Query("type"),
	})

	if err != nil {
		err = fiber.NewError(400, "cannot bring pokemons")
		c.Next(err)
		return
	}

	if len(res) == 0 {
		_ = c.JSON([]interface{}{})
	}

	c.JSON(res)
}

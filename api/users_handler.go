package api

import (
	"car-api/internal/middleware"
	"github.com/gofiber/fiber"
)

func (w *WebServices) CreateUserHandler(c *fiber.Ctx) {
	var cmd CreateUserCMD
	err := c.BodyParser(&cmd)

	res, err := w.Services.users.SaveUser(cmd)

	if err != nil {
		err = fiber.NewError(400, "cannot create user")
		c.Next(err)
		return
	}

	res.JWT = middleware.SignToken(w.tokenKey, res.ID)
	_ = c.JSON(res)

}

func (w *WebServices) WishListHandler(c *fiber.Ctx) {

	var cmd WishPokemonCMD
	_ = c.BodyParser(&cmd)

	bearer := c.Get("Authorization")
	userID := middleware.ExtractUserIDFromJWT(bearer, w.tokenKey)
	err := w.users.AddWishPokemon(userID, cmd.PokemonID, cmd.Comment)

	if err != nil {
		err = fiber.NewError(400, "cannot add to the wishlist")
		c.Next(err)
		return
	}

	_ = c.JSON(struct {
		Res string `json:"result"`
	}{
		Res: "pokemon added to the wishlist",
	})
}

func (w *WebServices) LoginHandler(c *fiber.Ctx) {
	var cmd LoginCMD
	err := c.BodyParser(&cmd)
	if err != nil {
		err = fiber.NewError(400, "cannot parse params")
		c.Next(err)
		return
	}

	id := w.users.Login(cmd)
	if id == "" {
		err = fiber.NewError(404, "user not found")
		c.Next(err)
		return
	}

	_ = c.JSON(struct {
		Token string `json:"token"`
	}{
		Token: middleware.SignToken(w.tokenKey, id),
	})
}

type LoginCMD struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type WishPokemonCMD struct {
	PokemonID string `json:"pokemon_id"`
	Comment   string `json:"comment"`
}

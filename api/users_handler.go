package api

import (
	"car-api/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func (w *WebServices) CreateUserHandler(c *fiber.Ctx) error {
	var cmd CreateUserCMD
	err := c.BodyParser(&cmd)

	res, err := w.Services.users.SaveUser(cmd)

	if err != nil {
		return fiber.NewError(400, "cannot create user")
	}

	res.JWT = middleware.SignToken(w.tokenKey, res.ID)
	return c.JSON(res)
}

func (w *WebServices) WishListHandler(c *fiber.Ctx) error {

	var cmd WishPokemonCMD
	_ = c.BodyParser(&cmd)

	bearer := c.Get("Authorization")
	userID := middleware.ExtractUserIDFromJWT(bearer, w.tokenKey)
	err := w.users.AddWishPokemon(userID, cmd.PokemonID, cmd.Comment)

	if err != nil {
		return fiber.NewError(400, "cannot add to the wishlist")
	}

	return c.JSON(struct {
		Res string `json:"result"`
	}{
		Res: "pokemon added to the wishlist",
	})
}

func (w *WebServices) LoginHandler(c *fiber.Ctx) error {
	var cmd LoginCMD
	err := c.BodyParser(&cmd)
	if err != nil {
		return fiber.NewError(400, "cannot parse params")
	}

	id := w.users.Login(cmd)
	if id == "" {
		return fiber.NewError(404, "user not found")
	}

	return c.JSON(struct {
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

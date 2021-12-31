package service

import (
	"car-api/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

// CreateUserHandler method for create a new access token.
// @Description Create a new User.
// @Summary create a new user
// @Tags Token
// @Accept json
// @Produce json
// @Param request body CreateUserCMD true "Create User Data"
// @Success 200 {string} status "ok"
// @Router /v1/create [post]
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

// LoginHandler method for create a new access token.
// @Description Login for get access token.
// @Summary Login for get access token
// @Tags Token
// @Accept json
// @Produce json
// @Param request body LoginCMD true "User Data"
// @Success 200 {string} status "ok"
// @Router /v1/login [post]
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
	token := middleware.SignToken(w.tokenKey, id)

	return c.JSON(fiber.Map{
		"error":        false,
		"msg":          nil,
		"access_token": token,
	})

	//return c.JSON(struct {
	//	Token string `json:"token"`
	//}{
	//	Token: middleware.SignToken(w.tokenKey, id),
	//})
}

type LoginCMD struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type WishPokemonCMD struct {
	PokemonID string `json:"pokemon_id"`
	Comment   string `json:"comment"`
}

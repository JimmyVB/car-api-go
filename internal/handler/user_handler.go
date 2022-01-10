package handler

import (
	user "car-api/internal/core/domain"
	"car-api/internal/core/ports"
	"car-api/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService ports.IUserService
	tokenKey    string
}

var _ ports.IUserHandler = (*UserHandler)(nil)

func NewUserHandler(userService ports.IUserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// Login method for create a new access token.
// @Description Login for get access token.
// @Summary Login for get access token
// @Tags Token
// @Accept json
// @Produce json
// @Param request body domain.User true "User Data"
// @Success 200 {string} status "ok"
// @Router /v1/login [post]
func (h *UserHandler) Login(c *fiber.Ctx) error {

	var user user.User
	err := c.BodyParser(&user)

	if err != nil {
		return fiber.NewError(400, "cannot parse params")
	}

	newUser, err := h.userService.Login(user)
	if err != nil {
		return fiber.NewError(404, "user not found")
	}

	token := middleware.SignToken(h.tokenKey, newUser)

	return c.JSON(fiber.Map{
		"error":        false,
		"msg":          nil,
		"access_token": token,
	})

}

// SaveUser method for create a new access token.
// @Description Create a new User.
// @Summary create a new user
// @Tags Token
// @Accept json
// @Produce json
// @Param request body domain.User true "Create User Data"
// @Success 200 {string} status "ok"
// @Router /v1/register [post]
func (h *UserHandler) SaveUser(c *fiber.Ctx) error {

	var user user.User
	err := c.BodyParser(&user)

	res, err := h.userService.SaveUser(user)

	if err != nil {
		return fiber.NewError(400, "cannot create user")
	}

	res.JWT = middleware.SignToken(h.tokenKey, res)

	return c.JSON(res)
}

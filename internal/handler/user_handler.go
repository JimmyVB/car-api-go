package handler

import (
	"car-api/internal/core/domain"
	enums "car-api/internal/core/emuns"
	"fmt"

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
// @Success 200 {object} domain.Message
// @Success 400 {object} domain.MessageError
// @Success 500 {object} domain.MessageError
// @Router /login [post]
func (h *UserHandler) Login(c *fiber.Ctx) error {

	var user domain.User
	err := c.BodyParser(&user)

	if err != nil {
		return fiber.NewError(400, "cannot parse params")
	}

	newUser, err := h.userService.Login(user)
	if err != nil {
		return fiber.NewError(404, "user not found")
	}

	token := middleware.SignToken(h.tokenKey, newUser)

	return c.JSON(domain.Message{
		Message: fmt.Sprintf("%s %s %s", enums.Logged, enums.Successfully, enums.User),
		Data:    token,
	})

}

// SaveUser method for create a new access token.
// @Description Create a new User.
// @Summary create a new user
// @Tags Token
// @Accept json
// @Produce json
// @Param request body domain.User true "Create User Data"
// @Success 200 {object} domain.Message
// @Success 400 {object} domain.MessageError
// @Success 500 {object} domain.MessageError
// @Router /register [post]
func (h *UserHandler) SaveUser(c *fiber.Ctx) error {

	var user domain.User
	err := c.BodyParser(&user)
	if err != nil {
		return fiber.NewError(400, "cannot parse params")
	}

	res, err := h.userService.SaveUser(user)

	if err != nil {
		return fiber.NewError(400, "cannot create user")
	}

	res.JWT = middleware.SignToken(h.tokenKey, res)

	return c.JSON(domain.Message{
		Message: fmt.Sprintf("%s %s %s", enums.Created, enums.Successfully, enums.User),
		Data:    res.JWT,
	})
}

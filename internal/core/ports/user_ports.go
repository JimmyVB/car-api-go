package ports

import (
	user "car-api/internal/core/domain"
	"github.com/gofiber/fiber/v2"
)

type IUserService interface {
	SaveUser(user user.User) (*user.UserResponse, error)
	Login(user user.User) (*user.UserResponse, error)
}

type IUserRepository interface {
	SaveUser(user user.User) (*user.UserResponse, error)
	Login(user user.User) (*user.UserResponse, error)
}

type IUserHandler interface {
	SaveUser(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type IServer interface {
	Initialize()
}

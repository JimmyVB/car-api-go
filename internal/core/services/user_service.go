package services

import (
	user "car-api/internal/core/domain"
	"car-api/internal/core/ports"
	"github.com/gofiber/fiber/v2"
)

type UserService struct {
	userRepository ports.IUserRepository
}

var _ ports.IUserService = (*UserService)(nil)

func NewUserService(repository ports.IUserRepository) *UserService {
	return &UserService{
		userRepository: repository,
	}
}

func (s *UserService) Login(user user.User) (*user.UserResponse, error) {
	res, err := s.userRepository.Login(user)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UserService) SaveUser(user user.User) (*user.UserResponse, error) {

	res, err := s.userRepository.SaveUser(user)
	if err != nil {
		return nil, fiber.NewError(400, "cannot create user")
	}

	return res, nil
}

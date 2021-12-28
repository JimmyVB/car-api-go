package api

import (
	"car-api/internal/database"
	"car-api/internal/logs"
	"github.com/gofiber/utils"
)

type CreateUserCMD struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserSummary struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	JWT      string `json:"token"`
}

type UserGateway interface {
	SaveUser(cmd CreateUserCMD) (*UserSummary, error)
	Login()
}

type UserService struct {
	*database.PostgresClient
}

func (us *UserService) SaveUser(cmd CreateUserCMD) (*UserSummary, error) {

	id := utils.UUID()

	_, err := us.Exec(CreateUserQuery(), id, cmd.Username, cmd.Password)
	if err != nil {
		logs.Error("cannot insert user " + err.Error())
		return nil, err
	}

	return &UserSummary{
		ID:       id,
		Username: cmd.Username,
		JWT:      "",
	}, nil
}

func (us *UserService) Login() {

}

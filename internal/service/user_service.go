package service

import (
	"car-api/internal/database"
	"car-api/internal/database/queries"
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
	Login(cmd LoginCMD) string
	AddWishPokemon(userId, pokemonId, comment string) error
}

type UserService struct {
	*database.PostgresClient
}

func (us *UserService) SaveUser(cmd CreateUserCMD) (*UserSummary, error) {

	id := utils.UUID()

	_, err := us.Exec(queries.CreateUserQuery(), id, cmd.Username, cmd.Password)
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

func (us *UserService) Login(cmd LoginCMD) string {
	var id string
	err := us.QueryRow(queries.GetLoginQuery(), cmd.Username, cmd.Password).Scan(&id)

	if err != nil {
		logs.Error(err.Error())
		return ""
	}
	return id
}

func (us *UserService) AddWishPokemon(userID, pokemonID, comment string) error {
	_, err := us.Exec(queries.GetAddWishPokemonQuery(), userID, pokemonID, comment)
	if err != nil {
		return err
	}
	return nil
}

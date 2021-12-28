package api

import "car-api/internal/database"

type Services struct {
	search PokemonSearch
	users  UserGateway
}

func NewServices() Services {
	client := database.NewPostgresClient()
	return Services{
		search: &PokemonService{client},
		users:  &UserService{client},
	}
}

type WebServices struct {
	Services
}

func start() *WebServices {
	return &WebServices{NewServices()}
}

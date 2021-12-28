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
	tokenKey string
}

func start(tokenKey string) *WebServices {
	return &WebServices{NewServices(), tokenKey}
}

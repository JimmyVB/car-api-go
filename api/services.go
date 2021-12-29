package api

import (
	"car-api/internal/database"
)

type Services struct {
	users UserGateway
	cars  CarGateway
}

func NewServices() Services {
	client := database.NewPostgresClient()
	return Services{
		users: &UserService{client},
		cars:  &CarService{PostgresClient: client},
	}
}

type WebServices struct {
	Services
	tokenKey string
}

func Start(tokenKey string) *WebServices {
	return &WebServices{NewServices(), tokenKey}
}

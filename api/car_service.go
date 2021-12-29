package api

import "car-api/internal/database"

type CreateCarCMD struct {
	ID    string `json:"id"`
	Mark  string `json:"mark"`
	Model string `json:"model"`
	Price uint   `json:"price"`
}

type CarService struct {
	*database.PostgresClient
}

type CarGateway interface {
	Save(id, mark, model string, price uint) error
}

func (us *CarService) Save(id, mark, model string, price uint) error {
	_, err := us.Exec(Insert(), id, mark, model, price)
	if err != nil {
		return err
	}
	return nil
}

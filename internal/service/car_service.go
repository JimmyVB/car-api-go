package service

import (
	"car-api/internal/database"
	"car-api/internal/database/queries"
	"car-api/internal/logs"
	"github.com/google/uuid"
)

type CreateCarCMD struct {
	ID    string `json:"id"`
	Marca string `json:"marca"`
	Model string `json:"model"`
	Price uint   `json:"price"`
}

type CarService struct {
	*database.PostgresClient
}

func (us *CarService) GetAll() ([]CreateCarCMD, error) {

	rows, err := us.Query(queries.GetAll())

	if err != nil {
		logs.Error(err.Error())
		return nil, err
	}

	var _cars []CreateCarCMD

	for rows.Next() {
		var carCMD CreateCarCMD
		err := rows.Scan(&carCMD.ID, &carCMD.Marca, &carCMD.Model, &carCMD.Price)
		if err != nil {
			logs.Error(err.Error())
			return nil, err
		}
		_cars = append(_cars, carCMD)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		logs.Error(err.Error())
		return nil, err
	}

	return _cars, err
}

func (us *CarService) GetOne(id uuid.UUID) (CreateCarCMD, error) {

	var _car CreateCarCMD
	err := us.QueryRow(queries.GetOne(), id).Scan(&_car.ID, &_car.Marca, &_car.Model, &_car.Price)

	if err != nil {
		logs.Error(err.Error())
		return CreateCarCMD{}, err
	}

	return _car, err
}

func (us *CarService) Update(id uuid.UUID, carCMD CreateCarCMD) (CreateCarCMD, error) {

	_, err := us.Exec(queries.Update(), id, carCMD.Marca, carCMD.Model, carCMD.Price)

	if err != nil {
		logs.Error(err.Error())
		return CreateCarCMD{}, err
	}

	return CreateCarCMD{}, err
}

type CarGateway interface {
	Save(id, mark, model string, price uint) error
	GetAll() ([]CreateCarCMD, error)
	GetOne(id uuid.UUID) (CreateCarCMD, error)
	Update(id uuid.UUID, carCMD CreateCarCMD) (CreateCarCMD, error)
}

func (us *CarService) Save(id, mark, model string, price uint) error {
	_, err := us.Exec(queries.Insert(), id, mark, model, price)
	if err != nil {
		return err
	}
	return nil
}

package repository

import (
	user "car-api/internal/core/domain"
	"car-api/internal/db/queries"
	"car-api/internal/logs"
	"database/sql"
)

type CarRepository struct {
	db *sql.DB
}

func NewCarRepository(db *sql.DB) *CarRepository {
	return &CarRepository{
		db: db,
	}
}

func (us *CarRepository) Save(car user.Car) error {

	_, err := us.db.Exec(queries.Insert(), car.Mark, car.Model, car.Price)
	if err != nil {
		logs.Error(err.Error())
		return err
	}
	return nil
}

func (us *CarRepository) GetOne(id string) (user.Car, error) {

	var _car user.Car
	err := us.db.QueryRow(queries.GetOne(), id).Scan(&_car.ID, &_car.Mark, &_car.Model, &_car.Price)

	if err != nil {
		logs.Error(err.Error())
		return user.Car{}, err
	}

	return _car, err
}

func (us *CarRepository) GetAll() ([]user.Car, error) {

	rows, err := us.db.Query(queries.GetAll())

	if err != nil {
		logs.Error(err.Error())
		return nil, err
	}

	var _cars []user.Car

	for rows.Next() {
		var carCMD user.Car
		err := rows.Scan(&carCMD.ID, &carCMD.Mark, &carCMD.Model, &carCMD.Price)
		if err != nil {
			logs.Error(err.Error())
			return nil, err
		}
		_cars = append(_cars, carCMD)
	}
	err = rows.Err()
	if err != nil {
		logs.Error(err.Error())
		return nil, err
	}

	return _cars, err
}

func (us *CarRepository) Update(id string, car user.Car) error {

	_, err := us.db.Exec(queries.Update(), id, car.Mark, car.Model, car.Price)
	if err != nil {
		logs.Error(err.Error())
		return err
	}
	return nil
}

func (us *CarRepository) Delete(id string) error {

	_, err := us.db.Exec(queries.Delete(), id)
	if err != nil {
		logs.Error(err.Error())
		return err
	}
	return nil
}

func (us *CarRepository) RentCar(carRent *user.CarRent) (*user.CarRent, error) {
	finalCarRent := user.CarRent{IdUser: 0, IdCar: 0}
	err := us.db.QueryRow(queries.GetCarStatus(), carRent.IdCar).Scan(&finalCarRent.ID, &finalCarRent.IdCar, &finalCarRent.IdUser, &finalCarRent.StartDate, &finalCarRent.EndDate)

	if err != nil {
		logs.Error(err.Error())
	}

	startYear := finalCarRent.StartDate.Year()
	endYear := finalCarRent.EndDate.Year()
	validRent := true
	if startYear != 1 && endYear != 1 &&
		((finalCarRent.StartDate.Before(carRent.StartDate) && finalCarRent.EndDate.After(carRent.EndDate)) ||
			(finalCarRent.StartDate.After(carRent.StartDate) && finalCarRent.StartDate.Before(carRent.EndDate)) ||
			(finalCarRent.EndDate.After(carRent.StartDate) && finalCarRent.EndDate.Before(carRent.EndDate))) {
		validRent = false
	}

	if validRent == true {
		_, err := us.db.Exec(queries.RentCar(), carRent.IdCar, carRent.IdUser, carRent.StartDate, carRent.EndDate)
		if err != nil {
			logs.Error(err.Error())
			return nil, err
		}
		return carRent, nil
	}
	return nil, nil
}

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

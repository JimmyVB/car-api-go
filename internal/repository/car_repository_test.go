package repository

import (
	"car-api/internal/core/domain"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_CourseRepository_Save_Succeed(t *testing.T) {

	id, mark, model, price := "3", "Mazda", "CX5", 25000

	car := domain.Car{ID: id, Mark: mark, Model: model, Price: uint(price)}

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"insert into cars (id, marca, model, price) values ($1, $2, $3, $4)").
		WithArgs(id, mark, model, price).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewCarRepository(db)
	err = repo.Save(car)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Empty(t, err)
}

func Test_CourseRepository_Save_Error(t *testing.T) {

	id, mark, model, price := "3", "Mazda", "CX5", 25000

	car := domain.Car{ID: id, Mark: mark, Model: model, Price: uint(price)}

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"insert into cars (id, marca, model, price) values ($1, $2, $3, $4)").
		WithArgs(id, mark, model, price).WillReturnError(err)

	repo := NewCarRepository(db)
	err = repo.Save(car)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
	assert.NotEmpty(t, err)
}

func TestCarRepository_GetOne_Succeed(t *testing.T) {

	id := "1"

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	carMockRows := sqlmock.NewRows([]string{"id", "mark", "model", "price"}).
		AddRow("1", "Toyota", "Hilux", "25000")

	sqlMock.ExpectQuery(
		"select id, marca, model, price from cars where id = $1").
		WithArgs(id).WillReturnRows(carMockRows)

	repo := NewCarRepository(db)
	_car, err := repo.GetOne("1")

	assert.NoError(t, err)
	require.Equal(t, _car.ID, "1")
}

func TestCarRepository_GetOne_Error(t *testing.T) {

	id := "1"

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery(
		"select id, marca, model, price from cars where id = $1").
		WithArgs(id).WillReturnError(err)

	repo := NewCarRepository(db)
	_car, err := repo.GetOne("1")

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
	assert.Empty(t, _car)
}

func TestCarRepository_GetAll_Succeed(t *testing.T) {

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	carMockRows := sqlmock.NewRows([]string{"id", "mark", "model", "price"}).
		AddRow("1", "Toyota", "Hilux", "25000").
		AddRow("2", "Ford", "Raptor", "45000")

	sqlMock.ExpectQuery(
		"select id, marca, model, price from cars").WillReturnRows(carMockRows)

	repo := NewCarRepository(db)
	_car, err := repo.GetAll()

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
	require.NotEmpty(t, _car)
}

func TestCarRepository_GetAll_Error(t *testing.T) {

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery(
		"select id, marca, model, price from cars").WillReturnError(err)

	repo := NewCarRepository(db)
	_car, err := repo.GetAll()

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
	require.Empty(t, _car)
}

func TestCarRepository_Update_Succeed(t *testing.T) {

	id, mark, model, price := "3", "Mazda", "CX9", 35000

	car := domain.Car{ID: id, Mark: mark, Model: model, Price: uint(price)}

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"update cars set marca = $2, model = $3, price = $4 where id = $1").
		WithArgs(id, mark, model, price).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewCarRepository(db)
	err = repo.Update(car.ID, car)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Empty(t, err)
}

func TestCarRepository_Update_Error(t *testing.T) {

	id, mark, model, price := "3", "Mazda", "CX9", 35000

	car := domain.Car{ID: id, Mark: mark, Model: model, Price: uint(price)}

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"update cars set marca = $2, model = $3, price = $4 where id = $1").
		WithArgs(id, mark, model, price).WillReturnError(err)

	repo := NewCarRepository(db)
	err = repo.Update(car.ID, car)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
	assert.NotEmpty(t, err)
}

func TestCarRepository_Delete_Succeed(t *testing.T) {
	id := "1"

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"delete from cars where id = $1").
		WithArgs(id).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewCarRepository(db)
	err = repo.Delete("1")

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
	assert.Empty(t, err)
}

func TestCarRepository_Delete_Error(t *testing.T) {
	id := "1"

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"delete from cars where id = $1").
		WithArgs(id).WillReturnError(err)

	repo := NewCarRepository(db)
	err = repo.Delete("1")

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
	assert.NotEmpty(t, err)
}

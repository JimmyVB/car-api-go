package repository

import (
	"car-api/internal/core/domain"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserRepository_Login_Succeed(t *testing.T) {

	user := domain.User{Username: "jimmy", Password: "jimmy123"}

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	userMockRows := sqlmock.NewRows([]string{"id"}).
		AddRow("1")

	sqlMock.ExpectQuery(
		"select id from users where username = $1 and password = $2").
		WithArgs(user.Username, user.Password).WillReturnRows(userMockRows)

	repo := NewUserRepository(db)
	res, err := repo.Login(user)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NotEmpty(t, res)
}

func TestUserRepository_Login_Error(t *testing.T) {

	user := domain.User{Username: "jimmy", Password: "jimmy123"}

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery(
		"select id from users where username = $1 and password = $2").
		WithArgs(user.Username, user.Password).WillReturnError(err)

	repo := NewUserRepository(db)
	res, err := repo.Login(user)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Empty(t, res)
}

func TestUserRepository_SaveUser_Succeed(t *testing.T) {

	user := domain.User{Username: "jimmy", Password: "jimmy123"}

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"insert into users (username, password) values ($1, $2)").
		WithArgs(user.Username, user.Password).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewUserRepository(db)
	_user, err := repo.SaveUser(user)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
	assert.NotEmpty(t, _user)
}

func TestUserRepository_SaveUser_Error(t *testing.T) {

	user := domain.User{}

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"insert into users (username, password) values ($1, $2)").
		WithArgs(user.Username, user.Password).WillReturnError(err)

	repo := NewUserRepository(db)
	_user, err := repo.SaveUser(user)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
	assert.Empty(t, _user)
}

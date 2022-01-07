package repository

import (
	user "car-api/internal/core/domain"
	userResponse "car-api/internal/core/domain"
	"car-api/internal/db/queries"
	"car-api/internal/logs"
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) Login(user user.User) string {
	var id string
	var newUser userResponse.UserResponse
	err := u.db.QueryRow(queries.GetLoginQuery(), user.Username, user.Password).Scan(&newUser.ID, &newUser.Username)

	if err != nil {
		logs.Error(err.Error())
		return ""
	}
	return id
}

func (u *UserRepository) SaveUser(user user.User) (*userResponse.UserResponse, error) {
	_, err := u.db.Exec(queries.CreateUserQuery(), user.Username, user.Password)
	if err != nil {
		logs.Error("cannot insert user " + err.Error())
		return nil, err
	}

	return &userResponse.UserResponse{
		Username: user.Username,
		JWT:      "",
	}, nil
}

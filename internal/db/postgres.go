package db

import (
	"car-api/internal/logs"
	"database/sql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PostgresClient struct {
	*sql.DB
}

func NewPostgresClient() *PostgresClient {
	db, err := sql.Open("postgres", "postgres://admin:admin@localhost:5432/crudgo?sslmode=disable")
	if err != nil {
		logs.Error("cannot create postgres client")
		panic(err)
	}

	err = db.Ping()

	return &PostgresClient{db}
}

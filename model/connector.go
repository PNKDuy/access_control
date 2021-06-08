package model

import (
	"database/sql"
)

func ConnectToPostgresAndReturnEnforcer() (*sql.DB, error){
	connectionString := "postgresql://postgres:86s25876@localhost:5432/access_control?sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

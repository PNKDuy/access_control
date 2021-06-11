package model

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

const (
	DB_DSN = "user=postgres password=86s25876 dbname=access_control sslmode=disable"
)

var Db *gorm.DB

func ConnectToPostgres() (*sql.DB, error){
	connectionString := "postgresql://postgres:86s25876@localhost:5432/access_control?sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func ConnectToPostgresWithGorm() {
	var err error
	Db, err = gorm.Open("postgres", DB_DSN)

	if err != nil {
		panic(err.Error())
	}

	// defer DB.Close()

	database := Db.DB()

	databaseError := database.Ping()

	if databaseError != nil {
		panic(databaseError.Error())
	}
}

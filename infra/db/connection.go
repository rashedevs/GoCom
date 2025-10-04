package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString() string {
	// user -> postgres
	// password -> redass
	// host -> localhost
	// port -> 5432
	// db name -> gocom
	return "user=postgres password=redass host=localhost port=5432 dbname=gocom sslmode=disable"
}

func NewConnection() (*sqlx.DB, error) {
	dbSource := GetConnectionString()
	dbCon, error := sqlx.Connect("postgres", dbSource)
	if error != nil {
		fmt.Println(error)
		return nil, error
	}
	return dbCon, nil
}

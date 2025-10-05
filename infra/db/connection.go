package db

import (
	"fmt"
	"gocom/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(cnf *config.DBConfig) string {
	connString := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s",
		cnf.User, cnf.Password, cnf.Host, cnf.Port, cnf.Name,
	)
	if !cnf.EnableSSLMode {
		connString += "sslmode=disable"
	}
	return connString
}

func NewConnection(cnf *config.DBConfig) (*sqlx.DB, error) {
	dbSource := GetConnectionString(cnf)
	dbCon, error := sqlx.Connect("postgres", dbSource)
	if error != nil {
		fmt.Println(error)
		return nil, error
	}
	return dbCon, nil
}

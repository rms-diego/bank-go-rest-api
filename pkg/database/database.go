package database

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/rms-diego/bank-go-rest-api/pkg/config"
)

var Db *sql.DB

func Connect(config *config.EnvironmentVariables) error {
	db, err := sql.Open("postgres", config.DatabaseUrl)

	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	Db = db
	return nil
}

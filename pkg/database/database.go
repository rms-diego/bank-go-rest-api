package database

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/rms-diego/bank-go-rest-api/pkg/config"
)

var Db *sql.DB

func initialize(db *sql.DB) error {
	createUserTable := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name 				VARCHAR(255) NOT NULL,
			last_name 	VARCHAR(255) NOT NULL,
			email 			VARCHAR(255) NOT NULL UNIQUE,
			tax_id 			VARCHAR(255) NOT NULL UNIQUE,
			birth_date 	DATE NOT NULL 
		);
	`

	_, err := db.Query(createUserTable)

	if err != nil {
		return err
	}

	return nil
}

func Connect(config *config.EnvironmentVariables) error {
	db, err := sql.Open("postgres", config.DatabaseUrl)

	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	initialize(db)

	Db = db
	return nil
}

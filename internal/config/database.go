package config

import (
	"database/sql"
	"os"
	"strings"
)

func Start() (*sql.DB, error) {

	dbUrl := os.Getenv("DB_URL")
	dbUrl = strings.TrimSpace(dbUrl)
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil, err
	}

	return db, nil
}

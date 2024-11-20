package config

import (
	"os"

	"github.com/pkg/errors"
)

func LoadDatabaseCredentials() (string, error) {
	database_dsn := os.Getenv("APP_DATABASE_DSN")

	if len(database_dsn) == 0 {
		return "", errors.New("database dsn not found")
	}

	return database_dsn, nil
}

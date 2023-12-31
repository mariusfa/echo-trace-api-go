package utils

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(dbConfig DbConfig, path string) error {
	if GetEnv("APP_ENV", "prod") == "dev" {
		return MigrateBase(dbConfig, path)
	}
	return nil
}

func MigrateBase(config DbConfig, path string) error {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.User, config.Password, config.Host, config.Port, config.DbName)

	basePath := fmt.Sprintf("file://%s/base", path)
	m, err := migrate.New(basePath, connectionString)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

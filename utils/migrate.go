package utils

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	defaultMigrationHost     = "localhost"
	defaultMigrationPort     = 5432
	defaultMigrationUser     = "postgres"
	defaultMigrationPassword = "password"
	defaultMigrationDBName   = "postgres"
)

type DbConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

func MigrateBase() {
	config := DbConfig{
		Host:     GetEnv("DB_HOST", defaultMigrationHost),
		Port:     GetEnvAsInt("DB_PORT", defaultMigrationPort),
		User:     GetEnv("DB_USER", defaultMigrationUser),
		Password: GetEnv("DB_PASSWORD", defaultMigrationPassword),
		DbName:   GetEnv("DB_NAME", defaultMigrationDBName),
	}

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.User, config.Password, config.Host, config.Port, config.DbName)

	m, err := migrate.New("file://migrations/base", connectionString)
	if err != nil {
		panic(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		panic(err)
	}
}
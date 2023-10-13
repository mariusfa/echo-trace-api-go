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

func MigrateBase() {
	host := GetEnv("DB_HOST", defaultMigrationHost)
	port := GetEnvAsInt("DB_PORT", defaultMigrationPort)
	user := GetEnv("DB_USER", defaultMigrationUser)
	password := GetEnv("DB_PASSWORD", defaultMigrationPassword)
	dbname := GetEnv("DB_NAME", defaultMigrationDBName)

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		user, password, host, port, dbname)

	m, err := migrate.New("file://migrations/base", connectionString)
	if err != nil {
		panic(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		panic(err)
	}
}
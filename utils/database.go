package utils

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/testcontainers/testcontainers-go"
)

type DbConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

func GetMigrationDbConfig() DbConfig {
	const (
		defaultMigrationHost     = "localhost"
		defaultMigrationPort     = 5432
		defaultMigrationUser     = "postgres"
		defaultMigrationPassword = "password"
		defaultMigrationDBName   = "postgres"
	)
	return DbConfig{
		Host:     GetEnv("DB_HOST", defaultMigrationHost),
		Port:     GetEnvAsInt("DB_PORT", defaultMigrationPort),
		User:     GetEnv("DB_USER", defaultMigrationUser),
		Password: GetEnv("DB_PASSWORD", defaultMigrationPassword),
		DbName:   GetEnv("DB_NAME", defaultMigrationDBName),
	}
}

func GetTestContainerMigrationDbConfig(container testcontainers.Container, ctx context.Context) (DbConfig, error) {
	host, err := container.Host(ctx)
	if err != nil {
		return DbConfig{}, err
	}
	port, err := container.MappedPort(ctx, "5432")
	if err != nil {
		return DbConfig{}, err
	}
	return DbConfig{
		Host:     host,
		Port:     port.Int(),
		User:     "test",
		Password: "test",
		DbName:   "test",
	}, nil
}

func GetTestContainerAppDbConfig(container testcontainers.Container, ctx context.Context) (DbConfig, error) {
	host, err := container.Host(ctx)
	if err != nil {
		return DbConfig{}, err
	}
	port, err := container.MappedPort(ctx, "5432")
	if err != nil {
		return DbConfig{}, err
	}
	return DbConfig{
		Host:     host,
		Port:     port.Int(),
		User:     "appuser",
		Password: "password",
		DbName:   "test",
	}, nil
}

func SetupAppDb(appDbConfig DbConfig) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		appDbConfig.Host, appDbConfig.Port, appDbConfig.User, appDbConfig.Password, appDbConfig.DbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}

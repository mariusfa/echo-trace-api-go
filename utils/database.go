package utils

import (
	"context"

	"github.com/testcontainers/testcontainers-go"
)

type DbConfig struct {
	Host            string
	Port            int
	User            string
	Password        string
	DbName          string
	migrationFolder string
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

func GetCustomMigrationDbConfig(host string, port int, user string, password string, dbName string) DbConfig {
	return DbConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DbName:   dbName,
	}
}

func GetTestContainerMigrationDbConfig(container testcontainers.Container, ctx context.Context, user string, password string, dbName string) (DbConfig, error) {
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
		User:     user,
		Password: password,
		DbName:   dbName,
	}, nil
}

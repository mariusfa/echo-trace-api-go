package biz_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"echo/biz"
	"echo/biz/domain"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestMain(m *testing.M) {
	ctx := context.Background()
	testContainer, err := createTestContainer(ctx)
	if err != nil {
		log.Fatalf("Failed to start container: %v", err)
	}
	defer testContainer.Terminate(ctx)

	// Run the tests
	code := m.Run()

	// Exit with the test code
	os.Exit(code)
}

func createTestContainer(ctx context.Context) (testcontainers.Container, error) {
	req := testcontainers.ContainerRequest{
		Image:        "postgres:13-alpine",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "test",
			"POSTGRES_PASSWORD": "test",
			"POSTGRES_DB":       "test",
		},
		WaitingFor: wait.ForLog("database system is ready to accept connections"),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	return container, nil
}

func TestUserRepository_Insert(t *testing.T) {
	repo := biz.NewUserRepository(db)

	user := domain.User{
		Name:           "testuser",
		HashedPassword: "testpass",
		ApiToken:       "testtoken",
	}
	err := repo.Insert(user)
	if err != nil {
		t.Errorf("Failed to insert user: %v", err)
	}

	// Check that the user was inserted correctly
	var result domain.User
	query := "SELECT id, name, hashed_password, api_token FROM echotraceschema.user WHERE name = $1"
	err = db.QueryRow(query, user.Name).Scan(&result.Id, &result.Name, &result.HashedPassword, &result.ApiToken)
	if err != nil {
		t.Errorf("Failed to get user: %v", err)
	}

	if result.Name != user.Name {
		t.Errorf("Unexpected user name: got %s, want %s", result.Name, user.Name)
	}
	if result.HashedPassword != user.HashedPassword {
		t.Errorf("Unexpected hashed password: got %s, want %s", result.HashedPassword, user.HashedPassword)
	}
	if result.ApiToken != user.ApiToken {
		t.Errorf("Unexpected API token: got %s, want %s", result.ApiToken, user.ApiToken)
	}
}

func TestUserRepository_GetByName(t *testing.T) {
	repo := biz.NewUserRepository(db)

	user := domain.User{
		Name:           "testuser",
		HashedPassword: "testpass",
		ApiToken:       "testtoken",
	}
	err := repo.Insert(user)
	if err != nil {
		t.Errorf("Failed to insert user: %v", err)
	}

	// Get the user by name
	result, err := repo.GetByName(user.Name)
	if err != nil {
		t.Errorf("Failed to get user: %v", err)
	}

	if result.Name != user.Name {
		t.Errorf("Unexpected user name: got %s, want %s", result.Name, user.Name)
	}
	if result.HashedPassword != user.HashedPassword {
		t.Errorf("Unexpected hashed password: got %s, want %s", result.HashedPassword, user.HashedPassword)
	}
	if result.ApiToken != user.ApiToken {
		t.Errorf("Unexpected API token: got %s, want %s", result.ApiToken, user.ApiToken)
	}
}

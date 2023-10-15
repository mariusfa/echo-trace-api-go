package biz_test

import (
	"context"
	"log"
	"os"
	"testing"

	"echo/biz"
	"echo/biz/domain"
	"echo/utils"
)

var userDbConfig utils.DbConfig

func TestMain(m *testing.M) {
	// Create the test container
	ctx := context.Background()
	testContainer, err := utils.CreateTestContainer(ctx)
	if err != nil {
		log.Fatalf("Failed to start container: %v", err)
	}
	defer testContainer.Terminate(ctx)

	// Get the migration db config
	migrationDbConfig, err := utils.GetTestContainerMigrationDbConfig(testContainer, ctx)
	if err != nil {
		log.Fatalf("Failed to get test container db config: %v", err)
	}
	// Do db migrations
	err = utils.MigrateBase(migrationDbConfig, "../migrations")

	// Get the app db config
	userDbConfig, err = utils.GetTestContainerAppDbConfig(testContainer, ctx)
	if err != nil {
		log.Fatalf("Failed to get test container app db config: %v", err)
	}

	// Run the tests
	code := m.Run()

	// Exit with the test code
	os.Exit(code)
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

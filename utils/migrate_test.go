package utils

import (
	"context"
	"testing"
)

func TestMigration(t *testing.T) {
	// Given
	ctx := context.Background()
	testContainer, err := CreateTestContainer(ctx)
	if err != nil {
		t.Fatalf("Failed to start container: %v", err)
	}
	defer testContainer.Terminate(ctx)
	migrationDbConfig, err := GetTestContainerMigrationDbConfig(testContainer, ctx, "test", "test", "test")
	if err != nil {
		t.Fatalf("Failed to get test container db config: %v", err)
	}

	// When
	err = MigrateBase(migrationDbConfig, "../migrations")

	// Then
	if err != nil {
		t.Fatalf("Failed to migrate: %v", err)
	}
}

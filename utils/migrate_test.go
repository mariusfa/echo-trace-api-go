package utils

import (
	"context"
	"testing"
)

func TestMigration(t *testing.T) {
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
	err = MigrateBase(migrationDbConfig, "../migrations")
	if err != nil {
		t.Fatalf("Failed to migrate: %v", err)
	}
}

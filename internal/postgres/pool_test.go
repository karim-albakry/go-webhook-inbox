package postgres

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestNewPool_ValidConnection_ReturnsUsablePool(t *testing.T) {
	// $env:TEST_DATABASE_URL = postgres://webhook_app:replace_with_your_local_password@localhost:5432/webhook_inbox?sslmode=disable
	db_url := os.Getenv("TEST_DATABASE_URL")
	if db_url == "" {
		t.Fatal("Environment variable [TEST_DATABASE_URL] is empty, make sure to provide it with database url")
	}

	ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
	defer cancel()

	pool, err := NewPool(ctx, db_url)

	if err != nil {
		t.Fatalf("Database Connection Error [error] %s", err.Error())
	}
	if pool == nil {
		t.Fatal("Pool is empty")
	}
	defer pool.Close()

	pingCtx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
	defer cancel()

	err = pool.Ping(pingCtx)
	if err != nil {
		t.Fatalf("Ping fail with error [Error] %s", err.Error())
	}
}

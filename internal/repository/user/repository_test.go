package user

import (
	"context"
	"os"
	"testing"

	"github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupTestDB(t *testing.T) *pgxpool.Pool {
	t.Helper()
	dbDSN := os.Getenv("APP_DATABASE_DSN")
	if len(dbDSN) == 0 {
		dbDSN = "postgres://testuser:testpass@localhost:5432/testdb?sslmode=disable"
	}
	pool, err := pgxpool.New(context.Background(), dbDSN)
	if err != nil {
		t.Fatalf("failed to connect to database: %vn", err)
	}
	return pool
}

func InitializeSchema(t *testing.T, pool *pgxpool.Pool) {
	t.Helper()
	_, err := pool.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY,
			first_name VARCHAR(255) NOT NULL,
			last_name VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL UNIQUE,
			created_at TIMESTAMP NOT NULL DEFAULT TIMEZONE('utc', now()),
			updated_at TIMESTAMP NOT NULL DEFAULT TIMEZONE('utc', now()),
			deleted_at TIMESTAMP
		);
	`)
	if err != nil {
		t.Fatalf("failed to initialize schema: %vn", err)
	}
}

func TeardownTestDB(t *testing.T, pool *pgxpool.Pool) {
	t.Helper()
	_, err := pool.Exec(context.Background(), "DROP TABLE IF EXISTS users;")
	if err != nil {
		t.Fatalf("failed to drop table: %vn", err)
	}
	pool.Close()
}

func TestCreateUser_OK(t *testing.T) {
	pool := SetupTestDB(t)
	InitializeSchema(t, pool)
	defer TeardownTestDB(t, pool)

	ctx := context.Background()
	tx, err := pool.Begin(ctx)
	if err != nil {
		t.Fatalf("failed to begin transaction: %vn", err)
	}
	defer tx.Rollback(ctx)

	repo := NewRepository(pool)

	user, err := repo.CreateUser(ctx, &model.CreateUserWithID{
		ID:        "12345678-1234-1234-1234-123456789012",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "oWw9S@example.com",
	})

	if err != nil {
		t.Fatalf("failed to create user: %vn", err)
	}

	if user.ID != "12345678-1234-1234-1234-123456789012" {
		t.Fatalf("unexpected user ID: %s", user.ID)
	}

	if user.FirstName != "John" {
		t.Fatalf("unexpected first name: %s", user.FirstName)
	}

	if user.LastName != "Doe" {
		t.Fatalf("unexpected last name: %s", user.LastName)
	}

	if user.Email != "oWw9S@example.com" {
		t.Fatalf("unexpected email: %s", user.Email)
	}
}

func TestFetchUserByID_UserNotFound(t *testing.T) {
	pool := SetupTestDB(t)
	InitializeSchema(t, pool)
	defer TeardownTestDB(t, pool)

	ctx := context.Background()
	tx, err := pool.Begin(ctx)
	if err != nil {
		t.Fatalf("failed to begin transaction: %vn", err)
	}
	defer tx.Rollback(ctx)

	repo := NewRepository(pool)

	_, err = repo.FetchUserById(ctx, "12345678-1234-1234-1234-123456789012")
	if err != model.ErrorUserNotFound {
		t.Fatalf("unexpected error: %v", err)
	}

}

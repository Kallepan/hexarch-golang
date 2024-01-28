package db

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
)

type Adapter struct {
	db *pgx.Conn
}

func getConnectionStringFromEnv() string {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbname)
}

func NewAdapter(ctx context.Context) (*Adapter, error) {
	// connect to db
	db, err := pgx.Connect(ctx, getConnectionStringFromEnv())
	if err != nil {
		slog.Error("db connection failed", "error", err)
		return nil, err
	}

	// ping
	err = db.Ping(ctx)
	if err != nil {
		slog.Error("db ping failed", "error", err)
		return nil, err
	}

	return &Adapter{db: db}, nil
}

func (a *Adapter) CloseDbConnection(ctx context.Context) error {
	err := a.db.Close(ctx)
	if err != nil {
		slog.Error("db connection close failed", "error", err)
		return err
	}

	return nil
}

func (a *Adapter) AddToHistory(ctx context.Context, x, y, result int32, operation string) error {
	queryString := "INSERT INTO arith_history (x, y, result, operation) VALUES ($1, $2, $3, $4)"
	params := []interface{}{x, y, result, operation}

	_, err := a.db.Exec(ctx, queryString, params...)
	if err != nil {
		slog.Error("db insert failed", "error", err)
		return err
	}

	return nil
}

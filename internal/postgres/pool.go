package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool(ctx context.Context, connectionString string) (*pgxpool.Pool, error) {
	connectCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(connectCtx, connectionString)
	if err != nil {
		return nil, err
	}
	err = pool.Ping(connectCtx)
	if err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}

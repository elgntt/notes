package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"task-manager/internal/config"
)

func OpenDB(ctx context.Context, cfg config.DBConfig) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig("")
	if err != nil {
		return nil, fmt.Errorf("OpenDB config parse: %w", err)
	}

	config.ConnConfig.Host = cfg.PgHost
	config.ConnConfig.Port = cfg.PgPort
	config.ConnConfig.Database = cfg.PgDatabase
	config.ConnConfig.User = cfg.PgUser
	config.ConnConfig.Password = cfg.PgPassword

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("OpenDB connect: %w", err)
	}

	if err = pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("OpenDB ping: %w", err)
	}

	return pool, nil
}

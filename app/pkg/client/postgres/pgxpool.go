package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type PostgresConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func NewPostgresConfig(username, password, host, port, database string) *PostgresConfig {
	return &PostgresConfig{
		Username: username,
		Password: password,
		Host:     host,
		Port:     port,
		Database: database,
	}
}

func NewClient(ctx context.Context, cfg *PostgresConfig, maxAttempts int, delay time.Duration) (pool *pgxpool.Pool) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	err := DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, delay)
		defer cancel()
		config, err := pgxpool.ParseConfig(dsn)
		if err != nil {
			return err
		}
		pool, err = pgxpool.NewWithConfig(ctx, config)
		if err != nil {
			return err
		}
		return nil

	}, maxAttempts, delay)
	if err != nil {
		return nil
	}

	return pool
}

func DoWithTries(fn func() error, maxAttempts int, delay time.Duration) (err error) {
	for maxAttempts > 0 {
		if err = fn(); err == nil {
			time.Sleep(delay)
			maxAttempts--
			continue
		}
		return nil
	}
	return
}

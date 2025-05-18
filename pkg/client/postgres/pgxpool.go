package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"log"
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
	err = Migrate(cfg)
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

func Migrate(cfg *PostgresConfig) error {
	goose.SetVerbose(true)
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.Username,
		cfg.Password, cfg.Host, cfg.Port,
		cfg.Database)
	sql, err := goose.OpenDBWithDriver("postgres", dbUrl)
	if err != nil {
		log.Fatalf("goose open: %w", err)
		return err
	}
	defer sql.Close()
	if err := goose.Up(sql, "./db/migrations"); err != nil {
		log.Fatalf("goose up: %v", err)
		return err
	}
	return nil
}

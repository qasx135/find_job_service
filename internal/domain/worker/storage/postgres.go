package storageworker

import (
	"context"
	"errors"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	worker "job_finder_service/internal/domain/worker/model"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

type Storage struct {
	client Client
}

func NewStorageWorker(client Client) *Storage {
	return &Storage{
		client: client,
	}
}

func (s *Storage) All(ctx context.Context) ([]worker.Worker, error) {
	from := sqlbuilder.Select(
		"id",
		"name",
		"surname",
		"patronymic").From("workers").String()

	rows, err := s.client.Query(ctx, from)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	list := make([]worker.Worker, 0)
	for rows.Next() {
		w := worker.Worker{}
		rows.Scan(&w.Id, &w.Name, &w.Surname, &w.Patronymic)
		list = append(list, w)
	}
	return list, nil
}

func (s *Storage) Create(ctx context.Context, worker *worker.Worker) error {
	q := `INSERT INTO workers (name,
surname,
patronymic) VALUES ($1, $2, $3)`
	if _, err := s.client.Exec(ctx, q, worker.Name, worker.Surname, worker.Patronymic); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return nil
			}
		}
	}
	return nil
}

package storageempl

import (
	"context"
	"errors"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"job_finder_service/internal/domain/employer/model"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

type Storage struct {
	client Client
}

func NewStorageEmp(client Client) *Storage {
	return &Storage{
		client: client,
	}
}

func (s *Storage) All(ctx context.Context) ([]employer.Employer, error) {
	from := sqlbuilder.Select("id", "name", "description").From("employers").String()

	rows, err := s.client.Query(ctx, from)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	list := make([]employer.Employer, 0)
	for rows.Next() {
		e := employer.Employer{}
		rows.Scan(&e.Id, &e.Name, &e.Description)
		list = append(list, e)
	}
	return list, nil
}

func (s *Storage) Create(ctx context.Context, employer *employer.Employer) error {
	q := `INSERT INTO employers (name, description) VALUES ($1, $2)`
	if _, err := s.client.Exec(ctx, q, employer.Name, employer.Description); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return nil
			}
		}
	}
	return nil
}

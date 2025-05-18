package storageresume

import (
	"context"
	"errors"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	resume "job_finder_service/internal/domain/resume/model"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

type Storage struct {
	client Client
}

func NewStorageResume(client Client) *Storage {
	return &Storage{
		client: client,
	}
}

func (s *Storage) All(ctx context.Context) ([]resume.Resume, error) {
	from := sqlbuilder.Select(
		"id",
		"worker_id",
		"about",
		"experience").From("resume").String()

	rows, err := s.client.Query(ctx, from)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	list := make([]resume.Resume, 0)
	for rows.Next() {
		r := resume.Resume{}
		rows.Scan(&r.Id, &r.WorkerId, &r.About, &r.Experience)
		list = append(list, r)
	}
	return list, nil
}

func (s *Storage) Create(ctx context.Context, resume *resume.Resume) error {
	q := `INSERT INTO resume (
    worker_id,
	about,
	experience) VALUES ($1, $2, $3)`
	if _, err := s.client.Exec(ctx, q, resume.WorkerId, resume.About, resume.Experience); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return nil
			}
		}
	}
	return nil
}

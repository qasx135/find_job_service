package storage

import (
	"context"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"job_finder_service/internal/domain/employer/model"
	"log"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

type Storage struct {
	client Client
}

func NewStorage(client Client) *Storage {
	return &Storage{
		client: client,
	}
}

func (s *Storage) All(ctx context.Context) ([]model.Employer, error) {
	from := sqlbuilder.Select("id", "name", "description").From("employers").String()

	rows, err := s.client.Query(ctx, from)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	list := make([]model.Employer, 0)
	for rows.Next() {
		e := model.Employer{}
		rows.Scan(&e.Id, &e.Name, &e.Description)
		list = append(list, e)
	}
	return list, nil
}

func (s *Storage) Create(ctx context.Context, employer *model.Employer) error {
	query := sqlbuilder.InsertInto("employers").Cols("name", "description").Values(employer.Name, employer.Description)
	_, err := s.client.Exec(ctx, query.String())
	if err != nil {
		log.Fatal("error inserting employer", err)
		return err
	}
	return nil
}

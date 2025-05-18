package storagejob

import (
	"context"
	"errors"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	job "job_finder_service/internal/domain/job/model"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

type Storage struct {
	client Client
}

func NewStorageJob(client Client) *Storage {
	return &Storage{
		client: client,
	}
}

func (s *Storage) All(ctx context.Context) ([]job.Job, error) {
	from := sqlbuilder.Select(
		"id",
		"header",
		"experience",
		"employment",
		"schedule",
		"work_format",
		"working_hours",
		"description",
		"employer_id").From("jobs").String()

	rows, err := s.client.Query(ctx, from)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	list := make([]job.Job, 0)
	for rows.Next() {
		j := job.Job{}
		rows.Scan(&j.Id, &j.Header, &j.Experience, &j.Employment,
			&j.Schedule, &j.WorkFormat, &j.WorkingHours, &j.Description,
			&j.EmployerId)
		list = append(list, j)
	}
	return list, nil
}

func (s *Storage) Create(ctx context.Context, job *job.Job) error {
	q := `INSERT INTO jobs (header,
experience,
employment,
schedule,
work_format,
working_hours,
description,
employer_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	if _, err := s.client.Exec(ctx, q, job.Header,
		job.Experience,
		job.Employment,
		job.Schedule,
		job.WorkFormat,
		job.WorkingHours,
		job.Description,
		job.EmployerId); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return nil
			}
		}
	}
	return nil
}

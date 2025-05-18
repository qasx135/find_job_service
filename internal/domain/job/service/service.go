package servicejob

import (
	"context"
	job "job_finder_service/internal/domain/job/model"
)

type PostgreRepo interface {
	All(ctx context.Context) ([]job.Job, error)
	Create(ctx context.Context, employer *job.Job) error
}

type Service struct {
	repo PostgreRepo
}

func NewServiceJob(repo PostgreRepo) *Service {
	return &Service{repo: repo}
}
func (s *Service) All(ctx context.Context) ([]job.Job, error) {
	return s.repo.All(ctx)
}

func (s *Service) Create(ctx context.Context, employer *job.Job) error {
	return s.repo.Create(ctx, employer)
}

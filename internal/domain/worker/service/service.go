package serviceworker

import (
	"context"
	worker "job_finder_service/internal/domain/worker/model"
)

type PostgreRepo interface {
	All(ctx context.Context) ([]worker.Worker, error)
	Create(ctx context.Context, worker *worker.Worker) error
}

type Service struct {
	repo PostgreRepo
}

func NewServiceWorker(repo PostgreRepo) *Service {
	return &Service{repo: repo}
}
func (s *Service) All(ctx context.Context) ([]worker.Worker, error) {
	return s.repo.All(ctx)
}

func (s *Service) Create(ctx context.Context, worker *worker.Worker) error {
	return s.repo.Create(ctx, worker)
}

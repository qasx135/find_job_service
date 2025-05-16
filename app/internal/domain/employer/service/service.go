package service

import (
	"context"
	"job_finder_service/internal/domain/employer/model"
)

type PostgreRepo interface {
	All(ctx context.Context) ([]model.Employer, error)
	Create(ctx context.Context, employer *model.Employer) error
}

type Service struct {
	repo PostgreRepo
}

func NewService(repo PostgreRepo) *Service {
	return &Service{repo: repo}
}
func (s *Service) All(ctx context.Context) ([]model.Employer, error) {
	return s.repo.All(ctx)
}

func (s *Service) Create(ctx context.Context, employer *model.Employer) error {
	return s.repo.Create(ctx, employer)
}

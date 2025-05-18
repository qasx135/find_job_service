package serviceempl

import (
	"context"
	"job_finder_service/internal/domain/employer/model"
)

type PostgreRepo interface {
	All(ctx context.Context) ([]employer.Employer, error)
	Create(ctx context.Context, employer *employer.Employer) error
}

type Service struct {
	repo PostgreRepo
}

func NewServiceEmpl(repo PostgreRepo) *Service {
	return &Service{repo: repo}
}
func (s *Service) All(ctx context.Context) ([]employer.Employer, error) {
	return s.repo.All(ctx)
}

func (s *Service) Create(ctx context.Context, employer *employer.Employer) error {
	return s.repo.Create(ctx, employer)
}

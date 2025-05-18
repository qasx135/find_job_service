package serviceresume

import (
	"context"
	resume "job_finder_service/internal/domain/resume/model"
)

type PostgreRepo interface {
	All(ctx context.Context) ([]resume.Resume, error)
	Create(ctx context.Context, employer *resume.Resume) error
}

type Service struct {
	repo PostgreRepo
}

func NewServiceResume(repo PostgreRepo) *Service {
	return &Service{repo: repo}
}
func (s *Service) All(ctx context.Context) ([]resume.Resume, error) {
	return s.repo.All(ctx)
}

func (s *Service) Create(ctx context.Context, resume *resume.Resume) error {
	return s.repo.Create(ctx, resume)
}

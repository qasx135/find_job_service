package handlers

import (
	"context"
	"job_finder_service/internal/domain/employer/model"
	"net/http"
)

type ServiceEmployer interface {
	Create(ctx context.Context, employer *model.Employer) error
	All(ctx context.Context) ([]model.Employer, error)
}

type Handler struct {
	Service ServiceEmployer
	ctx     context.Context
}

func NewHandler(ctx context.Context, employerService ServiceEmployer) *Handler {
	return &Handler{
		Service: employerService,
		ctx:     ctx,
	}
}

func (s *Handler) CreateEmployer(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}
func (s *Handler) AllEmployers(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

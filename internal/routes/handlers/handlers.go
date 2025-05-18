package handlers

import (
	"context"
	"github.com/go-chi/render"
	"job_finder_service/internal/domain/employer/model"
	job "job_finder_service/internal/domain/job/model"
	resume "job_finder_service/internal/domain/resume/model"
	worker "job_finder_service/internal/domain/worker/model"
	"log/slog"
	"net/http"
)

type ServiceEmployer interface {
	Create(ctx context.Context, employer *employer.Employer) error
	All(ctx context.Context) ([]employer.Employer, error)
}

type ServiceJob interface {
	Create(ctx context.Context, job *job.Job) error
	All(ctx context.Context) ([]job.Job, error)
}

type ServiceResume interface {
	Create(ctx context.Context, resume *resume.Resume) error
	All(ctx context.Context) ([]resume.Resume, error)
}

type ServiceWorker interface {
	Create(ctx context.Context, worker *worker.Worker) error
	All(ctx context.Context) ([]worker.Worker, error)
}

type Handler struct {
	ServiceEmpl   ServiceEmployer
	ServiceJob    ServiceJob
	ServiceRes    ServiceResume
	ServiceWorker ServiceWorker
	ctx           context.Context
}

func NewHandler(ctx context.Context,
	employerService ServiceEmployer,
	jobService ServiceJob,
	resumeService ServiceResume,
	workerService ServiceWorker) *Handler {
	return &Handler{
		ServiceEmpl:   employerService,
		ServiceJob:    jobService,
		ServiceRes:    resumeService,
		ServiceWorker: workerService,
		ctx:           ctx,
	}
}

func (s *Handler) CreateEmployer(w http.ResponseWriter, r *http.Request) {
	employer := &employer.Employer{}
	if err := render.DecodeJSON(r.Body, employer); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		slog.Error("Error decoding employer json", err)
	}
	err := s.ServiceEmpl.Create(s.ctx, employer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error("Error creating employer", err)
	}
}
func (s *Handler) AllEmployers(w http.ResponseWriter, r *http.Request) {
	allEmployers, err := s.ServiceEmpl.All(s.ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error("Error getting all employers from handlers", err)
	}
	render.JSON(w, r, allEmployers)
}

func (s *Handler) CreateJob(w http.ResponseWriter, r *http.Request) {
	job := &job.Job{}
	if err := render.DecodeJSON(r.Body, job); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		slog.Error("Error decoding job json", err)
	}
	err := s.ServiceJob.Create(s.ctx, job)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error("Error creating job", err)
	}
}
func (s *Handler) AllJobs(w http.ResponseWriter, r *http.Request) {
	allJobs, err := s.ServiceJob.All(s.ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error("Error getting all jobs from handlers", err)
	}
	render.JSON(w, r, allJobs)
}

func (s *Handler) CreateResume(w http.ResponseWriter, r *http.Request) {
	resume := &resume.Resume{}
	if err := render.DecodeJSON(r.Body, resume); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		slog.Error("Error decoding resume json", err)
	}
	err := s.ServiceRes.Create(s.ctx, resume)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error("Error creating resume", err)
	}
}
func (s *Handler) AllResume(w http.ResponseWriter, r *http.Request) {
	allResume, err := s.ServiceRes.All(s.ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error("Error getting all resume from handlers", err)
	}
	render.JSON(w, r, allResume)
}

func (s *Handler) CreateWorker(w http.ResponseWriter, r *http.Request) {
	worker := &worker.Worker{}
	if err := render.DecodeJSON(r.Body, worker); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		slog.Error("Error decoding worker json", err)
	}
	err := s.ServiceWorker.Create(s.ctx, worker)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error("Error creating worker", err)
	}
}
func (s *Handler) AllWorkers(w http.ResponseWriter, r *http.Request) {
	allWorkers, err := s.ServiceWorker.All(s.ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error("Error getting all workers from handlers", err)
	}
	render.JSON(w, r, allWorkers)
}

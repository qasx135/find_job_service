package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"job_finder_service/internal/routes/handlers"
)

type Router struct {
	Router  *chi.Mux
	Handler handlers.Handler
}

func NewRouter(h *handlers.Handler) *Router {
	cs := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	})

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(cs.Handler)
	r.Get("/employers", h.AllEmployers)

	return &Router{
		Router:  r,
		Handler: *h,
	}
}

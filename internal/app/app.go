package app

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"job_finder_service/internal/config"
	"job_finder_service/internal/domain/employer/service"
	"job_finder_service/internal/domain/employer/storage"
	"job_finder_service/internal/domain/job/service"
	storagejob "job_finder_service/internal/domain/job/storage"
	serviceresume "job_finder_service/internal/domain/resume/service"
	storageresume "job_finder_service/internal/domain/resume/storage"
	serviceworker "job_finder_service/internal/domain/worker/service"
	storageworker "job_finder_service/internal/domain/worker/storage"
	"job_finder_service/internal/routes"
	"job_finder_service/internal/routes/handlers"
	"job_finder_service/pkg/client/postgres"
	"log"
	"log/slog"
	"net"
	"net/http"
	"time"
)

type App struct {
	cfg        *config.Config
	router     *routes.Router
	httpServer *http.Server
	pgClient   *pgxpool.Pool
}

func NewApp(ctx context.Context, cfg *config.Config) (App, error) {

	slog.Info("router initializing")

	pgConfig := postgres.NewPostgresConfig(cfg.PostgreSQL.Username,
		cfg.PostgreSQL.Password,
		cfg.PostgreSQL.Host,
		cfg.PostgreSQL.Port,
		cfg.PostgreSQL.Database,
	)

	pgClient := postgres.NewClient(context.Background(), pgConfig, 5, 5*time.Second)

	newStorageEmpl := storageempl.NewStorageEmp(pgClient)
	newStorageJob := storagejob.NewStorageJob(pgClient)
	newStorageResume := storageresume.NewStorageResume(pgClient)
	newStorageWorker := storageworker.NewStorageWorker(pgClient)

	newServiceEmpl := serviceempl.NewServiceEmpl(newStorageEmpl)
	newServiceJob := servicejob.NewServiceJob(newStorageJob)
	newServiceResume := serviceresume.NewServiceResume(newStorageResume)
	newServiceWorker := serviceworker.NewServiceWorker(newStorageWorker)

	handler := handlers.NewHandler(ctx, newServiceEmpl, newServiceJob, newServiceResume, newServiceWorker)
	router := routes.NewRouter(handler)

	return App{cfg: cfg, router: router, pgClient: pgClient}, nil
}

func (app *App) Run() {
	app.StartHttpServer()
}

func (app *App) StartHttpServer() {
	slog.Info("starting http server")

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", app.cfg.Listen.Host, app.cfg.Listen.Port))
	slog.Info(fmt.Sprintf("binded host:port %s:%s", app.cfg.Listen.Host, app.cfg.Listen.Port))
	if err != nil {
		log.Fatal("error making listener: ", err)
	}

	app.httpServer = &http.Server{
		Handler:      app.router.Router,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
	}
	slog.Info("http server initialized and started")
	if err = app.httpServer.Serve(listener); err != nil {
		log.Fatal("error starting http server: ", err)
	}
}

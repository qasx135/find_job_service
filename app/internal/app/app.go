package app

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"job_finder_service/internal/config"
	"log"
	"log/slog"
	"net"
	"net/http"
	"time"
)

type App struct {
	cfg        *config.Config
	router     *httprouter.Router
	httpServer *http.Server
}

func NewApp(cfg *config.Config) (App, error) {

	slog.Info("router initializing")
	router := httprouter.New()

	router.Handler(http.MethodGet, "/swagger", http.RedirectHandler("/swagger/index.html", http.StatusMovedPermanently))
	router.Handler(http.MethodGet, "/swagger/*any", httpSwagger.WrapHandler)
	return App{cfg: cfg, router: router}, nil
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
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete},
	})
	handler := c.Handler(app.router)
	app.httpServer = &http.Server{
		Handler:      handler,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
	}
	slog.Info("http server initialized and started")
	if err = app.httpServer.Serve(listener); err != nil {
		log.Fatal("error starting http server: ", err)
	}
}

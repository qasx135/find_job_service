package main

import (
	app2 "job_finder_service/internal/app"
	"job_finder_service/internal/config"
	"log"
	"log/slog"
)

func main() {
	slog.Info("config initializing")
	cfg := config.GetInstance()

	app, err := app2.NewApp(cfg)
	if err != nil {
		log.Fatal("error initializing app: ", err)
	}
	slog.Info("application running")
	app.Run()
}

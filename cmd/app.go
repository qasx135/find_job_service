package main

import (
	"context"
	app2 "job_finder_service/internal/app"
	"job_finder_service/internal/config"
	"log"
	"log/slog"
)

func main() {
	ctx := context.Background()
	slog.Info("config initializing")
	cfg := config.GetInstance()

	app, err := app2.NewApp(ctx, cfg)
	if err != nil {
		log.Fatal("error initializing app: ", err)
	}
	slog.Info("application running")
	app.Run()
}

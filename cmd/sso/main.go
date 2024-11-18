package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/curtrika/UMetrika_server/internal/app"
	"github.com/curtrika/UMetrika_server/internal/config"
	"github.com/curtrika/UMetrika_server/internal/repository/postgres"
	dbGenerated "github.com/curtrika/UMetrika_server/internal/repository/postgres/generated"
	"github.com/curtrika/UMetrika_server/internal/storage"
)

// TODO: вынести в отдельный модуль
const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	ctx, cancel := context.WithCancel(context.Background())

	database, err := storage.DatabaseInit(cfg.DatabaseURL)
	if err != nil {
		panic(err)
	}

	testsRepo, err := postgres.New(ctx, cfg.DatabaseURL, &dbGenerated.ConverterImpl{})

	application := app.Init(ctx, log, cfg.GRPC.Port, cfg.TokenTTL, database, testsRepo, database)

	go func() {
		application.GRPCServer.MustRun()
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop
	cancel()
	application.GRPCServer.Stop()
	log.Info("Gracefully stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

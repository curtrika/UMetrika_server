package app

import (
	"database/sql"
	grpcapp "github.com/curtrika/UMetrika_server/internal/app/grpc"
	"github.com/curtrika/UMetrika_server/internal/services/auth"
	"github.com/curtrika/UMetrika_server/internal/storage"
	"log/slog"
	"time"
)

type App struct {
	GRPCServer *grpcapp.App
}

func Init(
	log *slog.Logger,
	grpcPort int,
	databaseURL string,
	tokenTTL time.Duration,
) *App {
	database, err := storage.DatabaseInit(databaseURL)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, database, database, database, tokenTTL)

	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}

// TODO: найти место получше
func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
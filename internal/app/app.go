package app

import (
	"context"
	"log/slog"
	"time"

	grpcapp "github.com/curtrika/UMetrika_server/internal/app/grpc"
	"github.com/curtrika/UMetrika_server/internal/domain/models"
	"github.com/curtrika/UMetrika_server/internal/services/admin_panel"
	"github.com/curtrika/UMetrika_server/internal/services/auth"
	"github.com/google/uuid"
)

type App struct {
	GRPCServer *grpcapp.App
}

type Repository interface {
	SaveUser(ctx context.Context, email string, passHash []byte) (uuid.UUID, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	// SaveApp(ctx context.Context, app models.App) (uuid.UUID, error)
	GetAppById(ctx context.Context, appID int32) (*models.App, error)

	CreateUser(ctx context.Context, user models.User) (*models.User, error)
	ReadUser(ctx context.Context, userID uuid.UUID) (*models.User, error)
	UpdateUser(ctx context.Context, user models.User) (*models.User, error)
	DeleteUser(ctx context.Context, userID uuid.UUID) error
}

func Init(
	ctx context.Context,
	log *slog.Logger,
	grpcPort int,
	tokenTTL time.Duration,
	repo Repository,
) *App {
	authService := auth.New(log, repo, repo, repo, tokenTTL)

	adminPanelService := admin_panel.New(log, repo)

	grpcApp := grpcapp.New(ctx, log, authService, adminPanelService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}

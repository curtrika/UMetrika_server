package app

import (
	"context"
	"log/slog"
	"time"

	"github.com/curtrika/UMetrika_server/internal/services/umetrika"

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
}

type testsRepo interface {
	CreateOwner(ctx context.Context, name string, email string, pass_hash []byte) (models.EducationOwner, error)
	CreateTest(ctx context.Context, testName string, description string, testType string, owner uuid.UUID) (models.EducationTest, error)
	GetOwner(ctx context.Context, ownerId uuid.UUID) (models.EducationOwner, error)
	GetTestsByOwnerId(ctx context.Context, ownerId uuid.UUID) ([]models.EducationTest, error)
	GetFullTestsByOwnerID(ctx context.Context, ownerId uuid.UUID) ([]models.EducationTestFull, error)
	InsertQuestionsToTest(ctx context.Context, questions []*models.QuestionAnswer) error
}

type schoolInfoProvider interface {
	GetTeacherDisciplinesAndClasses(ctx context.Context, teacherID uuid.UUID) ([]models.TeacherDiscipline, error)
}

func Init(
	ctx context.Context,
	log *slog.Logger,
	grpcPort int,
	tokenTTL time.Duration,
	repo Repository,
	tests testsRepo,
	schoolInfoProvider schoolInfoProvider,
) *App {
	authService := auth.New(log, repo, repo, repo, tokenTTL)

	adminPanelService := admin_panel.New(log, repo)

	umetrikaService := umetrika.New(log, tests, schoolInfoProvider)

	grpcApp := grpcapp.New(ctx, log, authService, adminPanelService, umetrikaService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}

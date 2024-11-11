package app

import (
	"context"
	"log/slog"
	"time"

	grpcapp "github.com/curtrika/UMetrika_server/internal/app/grpc"
	"github.com/curtrika/UMetrika_server/internal/domain/models"
	"github.com/curtrika/UMetrika_server/internal/services/admin_panel"
	"github.com/curtrika/UMetrika_server/internal/services/auth"
	storage "github.com/curtrika/UMetrika_server/internal/storage/sqlc_gen"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type App struct {
	GRPCServer *grpcapp.App
}

type Repository interface {
	SaveUser(ctx context.Context, email string, passHash []byte) (uuid.UUID, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	// SaveApp(ctx context.Context, app models.App) (uuid.UUID, error)

	// currently its easier to import storage models than to manualy copy-paste them to model packaeg
	GetAppById(ctx context.Context, appID int32) (*models.App, error)
	CreateAnswer(ctx context.Context, arg storage.CreateAnswerParams) (storage.Answer, error)
	CreateApp(ctx context.Context, arg storage.CreateAppParams) (storage.App, error)
	CreatePsychologicalPerformance(ctx context.Context, arg storage.CreatePsychologicalPerformanceParams) (storage.PsychologicalPerformance, error)
	CreatePsychologicalTest(ctx context.Context, arg storage.CreatePsychologicalTestParams) (storage.PsychologicalTest, error)
	CreatePsychologicalType(ctx context.Context, title string) (storage.PsychologicalType, error)
	CreateQuestion(ctx context.Context, arg storage.CreateQuestionParams) (storage.Question, error)
	CreateUser(ctx context.Context, arg storage.CreateUserParams) (storage.User, error)
	GetAnswer(ctx context.Context, id int32) (storage.Answer, error)
	GetApp(ctx context.Context, id int32) (storage.App, error)
	GetPsychologicalPerformance(ctx context.Context, id int32) (storage.PsychologicalPerformance, error)
	GetPsychologicalTest(ctx context.Context, id int32) (storage.PsychologicalTest, error)
	GetPsychologicalType(ctx context.Context, id int32) (storage.PsychologicalType, error)
	GetQuestion(ctx context.Context, id int32) (storage.Question, error)
	GetUser(ctx context.Context, id pgtype.UUID) (storage.User, error)
	ListAnswers(ctx context.Context) ([]storage.Answer, error)
	ListApps(ctx context.Context) ([]storage.App, error)
	ListPsychologicalPerformances(ctx context.Context) ([]storage.PsychologicalPerformance, error)
	ListPsychologicalTests(ctx context.Context) ([]storage.PsychologicalTest, error)
	ListPsychologicalTypes(ctx context.Context) ([]storage.PsychologicalType, error)
	ListQuestions(ctx context.Context) ([]storage.Question, error)
	ListUsers(ctx context.Context) ([]storage.User, error)
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

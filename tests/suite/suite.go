package suite

import (
	"context"
	"log/slog"
	"net"
	"os"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/curtrika/UMetrika_server/internal/app"
	"github.com/curtrika/UMetrika_server/internal/config"
	"github.com/curtrika/UMetrika_server/internal/domain/models"
	"github.com/curtrika/UMetrika_server/internal/storage"
	ssov1 "github.com/curtrika/UMetrika_server/pkg/proto/auth/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Suite struct {
	*testing.T                  // Потребуется для вызова методов *testing.T внутри Suite
	Cfg        *config.Config   // Конфигурация приложения
	AuthClient ssov1.AuthClient // Клиент для взаимодействия с gRPC-сервером
}

const (
	grpcHost = "localhost"
)

// same const from tests
const (
	emptyAppID = 0
	appID      = 1
	appSecret  = "test-secret"

	passDefaultLen = 10
)

func New(t *testing.T) (context.Context, *Suite) {
	t.Helper()

	// Automatically assign an available port
	cfg := config.Config{
		Env:         "local",
		DatabaseURL: os.Getenv("LOCAL_DB_TEST"),
		GRPC: config.GRPCConfig{
			Port:    44156, // use 0 for automatic port selection
			Timeout: 10 * time.Hour,
		},
		TokenTTL: time.Duration(99 * time.Hour),
	}

	ctx, cancelCtx := context.WithCancel(context.Background())

	// Initialize mock database and application
	mockDB := storage.NewMockDatabase()
	mockDB.SaveAppById(ctx, appID, models.App{
		ID:     appID,
		Name:   "test",
		Secret: appSecret,
	})
	application := app.Init(ctx, slog.Default(), cfg.GRPC.Port, cfg.TokenTTL, &mockDB)

	// Run GRPC server in a goroutine and ensure proper shutdown with WaitGroup
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := application.GRPCServer.Run(); err != nil {
			t.Fatalf("failed to start gRPC server: %v", err)
		}
	}()

	// Cleanup sequence: ensure server stops, then cancel the context
	t.Cleanup(func() {
		application.GRPCServer.Stop() // Gracefully shut down server
		wg.Wait()
		cancelCtx()
	})

	// Establish a client connection using the configured port
	cc, err := grpc.DialContext(ctx,
		grpcAddress(&cfg), // Note: ctx now used here
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("grpc server connection failed: %v", err)
	}

	return ctx, &Suite{
		T:          t,
		Cfg:        &cfg,
		AuthClient: ssov1.NewAuthClient(cc),
	}
}

func grpcAddress(cfg *config.Config) string {
	return net.JoinHostPort(grpcHost, strconv.Itoa(cfg.GRPC.Port))
}

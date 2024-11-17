package grpcapp

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	adminpanelgrpc "github.com/curtrika/UMetrika_server/internal/grpc/admin_panel"
	authgrpc "github.com/curtrika/UMetrika_server/internal/grpc/auth"
	umetrikagrpc "github.com/curtrika/UMetrika_server/internal/grpc/umetrika"
	"github.com/curtrika/UMetrika_server/internal/grpc/umetrika/generated"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

// MustRun runs gRPC server and panics if any error occurs.
func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

// Run runs gRPC server.
func (a *App) Run() error {
	const op = "grpcapp.Run"

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	a.log.Info("grpc server started", slog.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// Stop stops gRPC server.
func (a *App) Stop() {
	const op = "grpcapp.Stop"

	a.log.With(slog.String("op", op)).
		Info("stopping gRPC server", slog.Int("port", a.port))

	// Используем встроенный в gRPCServer механизм graceful shutdown
	a.gRPCServer.GracefulStop()
}

// New creates new gRPC server app.
func New(
	ctx context.Context,
	log *slog.Logger,
	authService authgrpc.Auth,
	adminPanelService adminpanelgrpc.AdminPanel,
	umetrikaService umetrikagrpc.UMetrika,
	port int,
) *App {
	loggingOpts := []logging.Option{ // позволяет логировать запросы/ответы сервера
		logging.WithLogOnEvents(
			logging.PayloadReceived, logging.PayloadSent,
		),
	}

	recoveryOpts := []recovery.Option{ // позволяет делать определенные действия при возникновении panic
		recovery.WithRecoveryHandler(func(p interface{}) (err error) {
			log.Error("Recovered from panic", slog.Any("panic", p))

			return status.Errorf(codes.Internal, "internal error")
		}),
	}

	// TODO: рассмотреть интерцепторы для трейсинга, метрик и алертов
	gRPCServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(recoveryOpts...),                       // восстанавливает и обрабатывает паники
		logging.UnaryServerInterceptor(InterceptorLogger(log), loggingOpts...), // логирует запросы/ответы сервера
	))

	// sso
	authgrpc.Register(gRPCServer, authService)

	// admin panel
	adminpanelgrpc.Register(gRPCServer, adminPanelService)
	go adminpanelgrpc.RunRest(ctx)

	umetrikagrpc.Register(gRPCServer, umetrikaService, &generated.ConverterImpl{})
	go umetrikagrpc.RunRest(ctx)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

// InterceptorLogger adapts slog logger to interceptor logger.
// This code is simple enough to be copied and not imported.
func InterceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}

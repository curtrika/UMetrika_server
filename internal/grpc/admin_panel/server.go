package adminpanelgrpc

import (
	"context"
	"github.com/curtrika/UMetrika_server/internal/converter"
	"github.com/curtrika/UMetrika_server/internal/domain/models"
	"log"
	"net/http"

	adminpanelv1 "github.com/curtrika/UMetrika_server/pkg/proto/admin_panel/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type serverAPI struct {
	adminpanelv1.UnimplementedAdminPanelServer
	adminPanel AdminPanel
	cvt        converter.GRPCConverter
}

type AdminPanel interface {
	// TODO: подумать над тем стоит ли выносить это по модулям. Наверное, да
	CreateUser(ctx context.Context, user models.User) (*models.User, error)
}

func Register(gRPCServer *grpc.Server, adminPanel AdminPanel) {
	adminpanelv1.RegisterAdminPanelServer(gRPCServer, &serverAPI{adminPanel: adminPanel})
}

func RunRest(ctx context.Context) {
	url := "localhost:44044" // rm hardcode
	grpcHandler := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := adminpanelv1.RegisterAdminPanelHandlerFromEndpoint(ctx, grpcHandler, url, opts)
	if err != nil {
		panic(err)
	}
	mux := http.NewServeMux()

	mux.Handle("/", grpcHandler)
	srv := http.Server{Addr: url, Handler: mux}

	log.Printf("server listening at 8081")

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	if err := srv.Shutdown(ctx); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}
}

func (s *serverAPI) Ping(ctx context.Context, req *adminpanelv1.PingMessage) (*adminpanelv1.PingMessage, error) {
	return req, nil
}

func (s *serverAPI) CreateUser(ctx context.Context, req *adminpanelv1.CreateUserRequest) (*adminpanelv1.CreateUserResponse, error) {
	userModel := s.cvt.UserToModel(req.UserData)
	newUser, err := s.adminPanel.CreateUser(ctx, *userModel)
	if err != nil {
		return nil, err
	}
	return &adminpanelv1.CreateUserResponse{
		NewUser: s.cvt.ModelToUser(newUser),
	}, nil
}

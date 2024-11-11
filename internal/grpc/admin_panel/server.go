package adminpanelgrpc

import (
	"context"
	"github.com/curtrika/UMetrika_server/internal/converter"
	"github.com/curtrika/UMetrika_server/internal/domain/models"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net/http"

	adminpanelv1 "github.com/curtrika/UMetrika_server/pkg/proto/admin_panel/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
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
	ReadUser(ctx context.Context, userID uuid.UUID) (*models.User, error)
	UpdateUser(ctx context.Context, userData models.User) (*models.User, error)
	DeleteUser(ctx context.Context, userID uuid.UUID) error
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
	//srv := http.Server{Addr: url, Handler: mux}

	log.Printf("server listening at 8081")

	withCors := cors.New(cors.Options{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"ACCEPT", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler(mux)

	if err := http.ListenAndServe(":8081", withCors); err != nil {
		panic(err)
	}
}

func (s *serverAPI) Ping(ctx context.Context, req *adminpanelv1.PingMessage) (*adminpanelv1.PingMessage, error) {
	return req, nil
}

func (s *serverAPI) CreateUser(ctx context.Context, req *adminpanelv1.CreateUserRequest) (*adminpanelv1.CreateUserResponse, error) {
	userModel := s.cvt.UserToModel(req.User)
	newUser, err := s.adminPanel.CreateUser(ctx, *userModel)
	if err != nil {
		return nil, err
	}
	return &adminpanelv1.CreateUserResponse{
		NewUser: s.cvt.ModelToUser(newUser),
	}, nil
}

func (s *serverAPI) ReadUser(ctx context.Context, req *adminpanelv1.ReadUserRequest) (*adminpanelv1.ReadUserResponse, error) {
	if req.UserId == "" {
		return nil, status.Error(codes.InvalidArgument, "missing user id")
	}

	parsedUserId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "missing user id")
	}

	user, err := s.adminPanel.ReadUser(ctx, parsedUserId)

	return &adminpanelv1.ReadUserResponse{
		User: s.cvt.ModelToUser(user),
	}, nil
}

func (s *serverAPI) UpdateUser(ctx context.Context, req *adminpanelv1.UpdateUserRequest) (*adminpanelv1.UpdateUserResponse, error) {
	//userModel := s.cvt.UserToModel(req.User)
	//newUser, err := s.adminPanel.CreateUser(ctx, *userModel)
	//if err != nil {
	//	return nil, err
	//}
	return nil, nil
}

func (s *serverAPI) DeleteUser(ctx context.Context, req *adminpanelv1.DeleteUserRequest) (*adminpanelv1.DeleteUserResponse, error) {
	//userModel := s.cvt.UserToModel(req.User)
	//newUser, err := s.adminPanel.CreateUser(ctx, *userModel)
	//if err != nil {
	//	return nil, err
	//}
	return nil, nil
}

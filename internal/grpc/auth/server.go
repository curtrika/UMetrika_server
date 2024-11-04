package server

import (
	"context"
	ssov1 "github.com/curtrika/UMetrika_server/pkg/proto/auth/v1"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type serverAPI struct {
	ssov1.UnimplementedAuthServer
	auth Auth
}

type Auth interface {
	Login(
		ctx context.Context,
		email string,
		password string,
		appID uuid.UUID,
	) (token string, err error)

	RegisterNewUser(
		ctx context.Context,
		email string,
		password string,
	) (userID uuid.UUID, err error)
}

func Register(gRPCServer *grpc.Server, auth Auth) {
	ssov1.RegisterAuthServer(gRPCServer, &serverAPI{auth: auth})
}

func (s *serverAPI) Login(
	ctx context.Context,
	in *ssov1.LoginRequest,
) (*ssov1.LoginResponse, error) {
	// TODO не забыть реализовать
}

func (s *serverAPI) Register(
	ctx context.Context,
	in *ssov1.RegisterRequest,
) (*ssov1.RegisterResponse, error) {
	// TODO не забыть реализовать
}

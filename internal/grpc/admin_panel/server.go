package adminpanelgrpc

import (
	"context"
	adminpanelv1 "github.com/curtrika/UMetrika_server/pkg/proto/admin_panel/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

type serverAPI struct {
	adminpanelv1.UnimplementedAdminPanelServer
	adminPanel AdminPanel
}

type AdminPanel interface {
}

func Register(gRPCServer *grpc.Server, adminPanel AdminPanel) {
	adminpanelv1.RegisterAdminPanelServer(gRPCServer, &serverAPI{adminPanel: adminPanel})
}

func RunRest() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := adminpanelv1.RegisterAdminPanelHandlerFromEndpoint(ctx, mux, "localhost:44044", opts)
	if err != nil {
		panic(err)
	}

	log.Printf("server listening at 8081")

	if err := http.ListenAndServe(":8081", mux); err != nil {
		panic(err)
	}
}

func (s *serverAPI) Ping(ctx context.Context, req *adminpanelv1.PingMessage) (*adminpanelv1.PingMessage, error) {
	return req, nil
}
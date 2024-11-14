package umetrikagrpc

import (
	"context"
	umetrikav1 "github.com/curtrika/UMetrika_server/pkg/proto/umetrika/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

type serverAPI struct {
	umetrikav1.UnimplementedUMetrikaServer
	umetrika UMetrika
}

type UMetrika interface{}

func Register(gRPCServer *grpc.Server, umetrika UMetrika) {
	umetrikav1.RegisterUMetrikaServer(gRPCServer, &serverAPI{umetrika: umetrika})
}

func RunRest(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := umetrikav1.RegisterUMetrikaHandlerFromEndpoint(ctx, mux, "localhost:44044", opts)
	if err != nil {
		panic(err)
	}

	log.Printf("server listening at 8082")

	withCors := cors.New(cors.Options{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"ACCEPT", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler(mux)

	if err := http.ListenAndServe(":8082", withCors); err != nil {
		panic(err)
	}
}

func (s *serverAPI) Ping(ctx context.Context, _ *umetrikav1.EmptyMessage) (*umetrikav1.PingMessage, error) {
	res := &umetrikav1.PingMessage{
		Message: "pong",
	}
	return res, nil
}

package umetrikagrpc

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/curtrika/UMetrika_server/internal/domain/models"
	umetrikav1 "github.com/curtrika/UMetrika_server/pkg/proto/umetrika/v1"
	"github.com/google/uuid"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type serverAPI struct {
	umetrikav1.UnimplementedUMetrikaServer
	umetrika  UMetrika
	converter Converter
}

type testsBackend interface {
	CreateOwner(ctx context.Context, name string, email string, pass_hash []byte) (models.EducationOwner, error)
	CreateTest(ctx context.Context, testName string, description string, testType string) (models.EducationTest, error)
	GetOwner(ctx context.Context, ownerId uuid.UUID) (models.EducationOwner, error)
	GetTestsByOwnerId(ctx context.Context, ownerId uuid.UUID) ([]models.EducationTest, error)
}

type UMetrika interface {
	testsBackend
}

func Register(gRPCServer *grpc.Server, umetrika UMetrika, converter Converter) {
	umetrikav1.RegisterUMetrikaServer(gRPCServer,
		&serverAPI{
			umetrika:  umetrika,
			converter: converter,
		},
	)
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

// POST api/v1/umetrica/owner
func (s *serverAPI) CreateOwner(ctx context.Context, req *umetrikav1.OwnerPost) (*umetrikav1.OwnerResult, error) {
	if req.Email == "" || req.Password == "" {
		return nil, fmt.Errorf("not enough data")
	}
	res, err := s.umetrika.CreateOwner(ctx, req.OwnerName, req.Email, []byte(req.Password))
	if err != nil {
		return nil, fmt.Errorf("error with creating owner: %w", err)
	}
	return s.converter.OwnerModelToProto(&res), nil
}

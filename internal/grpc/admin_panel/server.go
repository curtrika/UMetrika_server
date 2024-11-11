package adminpanelgrpc

import (
	"context"
	"log"
	"net/http"

	adminpanelv1 "github.com/curtrika/UMetrika_server/pkg/proto/admin_panel/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc/credentials/insecure"
)

type serverAPI struct {
	adminpanelv1.UnimplementedAdminPanelServer
	adminPanel AdminPanel
}

type AdminPanel interface{}

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

func (s *serverAPI) CreateDiscipline(context.Context, *adminpanelv1.CreateDisciplineRequest) (*adminpanelv1.CreateDisciplineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDiscipline not implemented")
}

func (s *serverAPI) CreatePsychologicalType(context.Context, *adminpanelv1.CreatePsychologicalTypeRequest) (*adminpanelv1.PsychologicalTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePsychologicalType not implemented")
}

func (s *serverAPI) GetPsychologicalType(context.Context, *adminpanelv1.GetPsychologicalTypeRequest) (*adminpanelv1.PsychologicalTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPsychologicalType not implemented")
}

func (s *serverAPI) ListPsychologicalTypes(context.Context, *adminpanelv1.ListPsychologicalTypesRequest) (*adminpanelv1.ListPsychologicalTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPsychologicalTypes not implemented")
}

func (s *serverAPI) CreatePsychologicalTest(context.Context, *adminpanelv1.CreatePsychologicalTestRequest) (*adminpanelv1.PsychologicalTestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePsychologicalTest not implemented")
}

func (s *serverAPI) GetPsychologicalTest(context.Context, *adminpanelv1.GetPsychologicalTestRequest) (*adminpanelv1.PsychologicalTestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPsychologicalTest not implemented")
}

func (s *serverAPI) ListPsychologicalTests(context.Context, *adminpanelv1.ListPsychologicalTestsRequest) (*adminpanelv1.ListPsychologicalTestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPsychologicalTests not implemented")
}

func (s *serverAPI) CreateQuestion(context.Context, *adminpanelv1.CreateQuestionRequest) (*adminpanelv1.QuestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQuestion not implemented")
}

func (s *serverAPI) GetQuestion(context.Context, *adminpanelv1.GetQuestionRequest) (*adminpanelv1.QuestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuestion not implemented")
}

func (s *serverAPI) ListQuestions(context.Context, *adminpanelv1.ListQuestionsRequest) (*adminpanelv1.ListQuestionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQuestions not implemented")
}

func (s *serverAPI) CreateAnswer(context.Context, *adminpanelv1.CreateAnswerRequest) (*adminpanelv1.AnswerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAnswer not implemented")
}

func (s *serverAPI) GetAnswer(context.Context, *adminpanelv1.GetAnswerRequest) (*adminpanelv1.AnswerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnswer not implemented")
}

func (s *serverAPI) ListAnswers(context.Context, *adminpanelv1.ListAnswersRequest) (*adminpanelv1.ListAnswersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAnswers not implemented")
}

func (s *serverAPI) CreatePsychologicalPerformance(context.Context, *adminpanelv1.CreatePsychologicalPerformanceRequest) (*adminpanelv1.PsychologicalPerformanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePsychologicalPerformance not implemented")
}

func (s *serverAPI) GetPsychologicalPerformance(context.Context, *adminpanelv1.GetPsychologicalPerformanceRequest) (*adminpanelv1.PsychologicalPerformanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPsychologicalPerformance not implemented")
}

func (s *serverAPI) ListPsychologicalPerformances(context.Context, *adminpanelv1.ListPsychologicalPerformancesRequest) (*adminpanelv1.ListPsychologicalPerformancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPsychologicalPerformances not implemented")
}

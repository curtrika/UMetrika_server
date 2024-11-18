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
	CreateTest(ctx context.Context, testName string, description string, testType string, owner uuid.UUID) (models.EducationTest, error)
	GetOwner(ctx context.Context, ownerId uuid.UUID) (models.EducationOwner, error)
	GetTestsByOwnerId(ctx context.Context, ownerId uuid.UUID) ([]models.EducationTest, error)
	InsertQuestionsToTest(ctx context.Context, questions []*models.QuestionAnswer) error
	GetFullTestsByOwnerId(ctx context.Context, ownerId uuid.UUID) ([]models.EducationTestFull, error)
}

type umetrikaProvider interface {
	GetTeacherDisciplinesAndClasses(ctx context.Context, teacherID uuid.UUID) ([]models.TeacherDiscipline, error)
}

type UMetrika interface {
	testsBackend
	umetrikaProvider
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

func (s *serverAPI) AddNewTest(ctx context.Context, req *umetrikav1.TestPost) (*umetrikav1.TestResult, error) {
	if req.TestName == "" || len(req.Questions) == 0 || req.OwnerId == "" {
		return nil, fmt.Errorf("not enough data")
	}
	u, err := StringToUUID(req.OwnerId)
	if err != nil {
		return nil, fmt.Errorf("wrong id format")
	}
	questions, err := s.converter.QuestionDTOsProtoToModel(req.Questions)
	if err != nil {
		return nil, fmt.Errorf("errors with questions conversion")
	}

	test, err := s.umetrika.CreateTest(ctx, req.TestName, req.Description, req.TestType, u)
	if err != nil {
		return nil, fmt.Errorf("error inserting test: %w", err)
	}
	fmt.Println(test)

	var questionAnswer []*models.QuestionAnswer
	for i, question := range questions {
		question.TestID = test.TestID
		questionAnswer = append(questionAnswer, &models.QuestionAnswer{
			TestID:    test.TestID,
			Questions: *question,
		})
		answers, err := s.converter.AnswerDTOsProtoToModel(req.Questions[i].Answers)
		if err != nil {
			return nil, fmt.Errorf("errors with questions conversion")
		}
		for _, answer := range answers {
			questionAnswer[i].Answers = append(questionAnswer[i].Answers, *answer)
		}
	}
	err = s.umetrika.InsertQuestionsToTest(ctx, questionAnswer)
	if err != nil {
		return nil, err
	}

	return s.converter.TestModelToProto(&test)
}

func (s *serverAPI) GetFullTestByOwnerId(ctx context.Context, req *umetrikav1.TestOwnerGet) (*umetrikav1.TestsGet, error) {
	u, err := StringToUUID(req.OwnerId)
	if err != nil {
		return nil, fmt.Errorf("wrong uuid format: %w", err)
	}
	fmt.Println(u.String())
	test, err := s.umetrika.GetFullTestsByOwnerId(ctx, u)
	if err != nil {
		return nil, err
	}
	return ConvertFullTestToProto(test), nil
}

func (s *serverAPI) GetTeacherDisciplinesAndClasses(ctx context.Context, req *umetrikav1.GetTeacherDisciplinesAndClassesRequest) (*umetrikav1.GetTeacherDisciplinesAndClassesResponse, error) {
	uid, err := uuid.Parse(req.TeacherId)
	if err != nil {
		return nil, fmt.Errorf("wrong id format")
	}

	teacherDisciplines, err := s.umetrika.GetTeacherDisciplinesAndClasses(ctx, uid)
	if err != nil {
		return nil, err
	}

	// TODO: подумать как получше сделать
	res := umetrikav1.GetTeacherDisciplinesAndClassesResponse{}
	for _, elem := range teacherDisciplines {
		pbElem, err := s.converter.TeacherDisciplineToProto(elem)
		if err != nil {
			return nil, err
		}
		res.TeacherDiscipline = append(res.TeacherDiscipline, pbElem)
	}

	return &res, nil
}

func ConvertFullTestToProto(tests []models.EducationTestFull) *umetrikav1.TestsGet {
	var protoTests []*umetrikav1.TestGet

	for _, test := range tests {
		// Convert questions
		var protoQuestions []*umetrikav1.QuestionGetDTO
		for _, question := range test.Questions {
			// Convert answers
			var protoAnswers []*umetrikav1.AnswerGetDTO
			for _, answer := range question.Answers {
				protoAnswers = append(protoAnswers, &umetrikav1.AnswerGetDTO{
					AnswerId:    answer.AnswerID.String(),
					AnswerText:  answer.AnswerText,
					AnswerOrder: answer.AnswerOrder,
				})
			}

			// Add converted question
			protoQuestions = append(protoQuestions, &umetrikav1.QuestionGetDTO{
				QuestionId:    question.QuestionID.String(),
				QuestionText:  question.QuestionText,
				QuestionOrder: question.QuestionOrder,
				QuestionType:  question.QuestionType,
				Answers:       protoAnswers,
			})
			fmt.Println(protoAnswers)
		}

		// Add converted test
		protoTests = append(protoTests, &umetrikav1.TestGet{
			TestId:      test.TestID.String(),
			TestName:    test.TestName,
			OwnerId:     test.OwnerID.String(),
			Description: test.Description,
			TestType:    test.TestType,
			Questions:   protoQuestions,
		})
	}

	return &umetrikav1.TestsGet{
		Tests: protoTests,
	}
}

// func (UnimplementedUMetrikaServer) AddNewTest(context.Context, *TestPost) (*TestResult, error) {

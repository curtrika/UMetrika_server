package postgres

import (
	"fmt"

	"github.com/curtrika/UMetrika_server/internal/domain/models"
	postgres "github.com/curtrika/UMetrika_server/internal/repository/postgres/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/net/context"
)

type TestRepository struct {
	queries *postgres.Queries
	mapper  Converter
}

func New(ctx context.Context, databaseURL string, mapper Converter) (*TestRepository, error) {
	conn, err := pgx.Connect(ctx, databaseURL)
	if err != nil {
		return nil, err
	}
	return &TestRepository{
		queries: postgres.New(conn),
		mapper:  mapper,
	}, nil
}

func (q *TestRepository) CreateOwner(ctx context.Context, name, email string, pass_hash []byte) (models.EducationOwner, error) {
	var params postgres.CreateOwnerParams
	params.Email = email
	params.OwnerName = name
	params.PassHash = pass_hash
	owner, err := q.queries.CreateOwner(ctx, params)
	if err != nil {
		return models.EducationOwner{}, nil
	}
	return q.mapper.OwnerDBToModel(owner), nil
}

func (q *TestRepository) GetOwner(ctx context.Context, id uuid.UUID) (models.EducationOwner, error) {
	var pgOwnerID pgtype.UUID
	if err := pgOwnerID.Scan(id); err != nil {
		return models.EducationOwner{}, fmt.Errorf("wrong uuid")
	}
	owner, err := q.queries.GetOwner(ctx, pgOwnerID)
	if err != nil {
		return models.EducationOwner{}, fmt.Errorf("error while getting owner by id: %w", err)
	}
	return q.mapper.OwnerDBToModel(owner), nil
}

func (q *TestRepository) CreateTest(ctx context.Context, testName, description, testType string) (models.EducationTest, error) {
	var params postgres.CreateTestParams
	params.TestName = testName
	if err := params.Description.Scan(description); err != nil {
		return models.EducationTest{}, fmt.Errorf("err while scanning description: %w", err)
	}
	if err := params.TestType.Scan(testType); err != nil {
		return models.EducationTest{}, fmt.Errorf("err while scanning testType: %w", err)
	}
	test, err := q.queries.CreateTest(ctx, params)
	if err != nil {
		return models.EducationTest{}, fmt.Errorf("error while creating test: %w", err)
	}
	return q.mapper.TestDBToModel(test), nil
}

func (q *TestRepository) GetTestsByOwnerId(ctx context.Context, ownerID uuid.UUID) ([]models.EducationTest, error) {
	var pgOwnerID pgtype.UUID
	if err := pgOwnerID.Scan(ownerID); err != nil {
		return nil, fmt.Errorf("wrong uuid")
	}
	tests, err := q.queries.GetTestsByOwnerId(ctx, pgOwnerID)
	if err != nil {
		return nil, fmt.Errorf("error while getting tests by owner id: %w", err)
	}
	return q.mapper.TestsDBToModel(tests), nil
}

// func (q *TestRepository) CreateAnswer(ctx context.Context, nextAnswerID uuid.UUID, text string) (models.Answer, error) {
// 	params := postgres.CreateAnswerParams{
// 		Title: text,
// 	}
// 	if err := params.NextAnswerID.Scan(nextAnswerID); err != nil {
// 		return models.Answer{}, fmt.Errorf("wrong nextAnswerId")
// 	}
// 	// return models.Answer{}, nil
// 	answer, err := q.queries.CreateAnswer(ctx, params)
// 	if err != nil {
// 		return models.Answer{}, fmt.Errorf("error at db: %w", err)
// 	}
// 	return q.mapper.AnswerDBToModel(answer), nil
// }
//
// func (q *TestRepository) CreatePsychologicalTest(ctx context.Context, ownerID uuid.UUID, title string) (models.PsychologicalTest, error) {
// 	params := postgres.CreatePsychologicalTestParams{
// 		Title:   title,
// 		OwnerID: 0,
// 	}
// 	test, err := q.queries.CreatePsychologicalTest(ctx, params)
// 	if err != nil {
// 		return models.PsychologicalTest{}, fmt.Errorf("error at db: %w", err)
// 	}
// 	return q.mapper.PsychologicalTestDBToModel(test), nil
// }

// func (q *TestRepository) CreateApp(ctx context.Context, arg CreateAppParams) (App, error)
// func (q *TestRepository) CreatePsychologicalPerformance(ctx context.Context, arg CreatePsychologicalPerformanceParams) (PsychologicalPerformance, error)
// func (q *TestRepository) CreatePsychologicalType(ctx context.Context, title string) (PsychologicalType, error)
// func (q *TestRepository) CreateQuestion(ctx context.Context, arg CreateQuestionParams) (Question, error)
// func (q *TestRepository) CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
// func (q *TestRepository) GetAnswer(ctx context.Context, id int32) (Answer, error)
// func (q *TestRepository) GetApp(ctx context.Context, id int32) (App, error)
// func (q *TestRepository) GetPsychologicalPerformance(ctx context.Context, id int32) (PsychologicalPerformance, error)
// func (q *TestRepository) GetPsychologicalTest(ctx context.Context, id int32) (PsychologicalTest, error)
// func (q *TestRepository) GetPsychologicalType(ctx context.Context, id int32) (PsychologicalType, error)
// func (q *TestRepository) GetQuestion(ctx context.Context, id int32) (Question, error)
// func (q *TestRepository) GetUser(ctx context.Context, id pgtype.UUID) (User, error)
// func (q *TestRepository) ListAnswers(ctx context.Context) ([]Answer, error)
// func (q *TestRepository) ListApps(ctx context.Context) ([]App, error)
// func (q *TestRepository) ListPsychologicalPerformances(ctx context.Context) ([]PsychologicalPerformance, error)
// func (q *TestRepository) ListPsychologicalTests(ctx context.Context) ([]PsychologicalTest, error)
// func (q *TestRepository) ListPsychologicalTypes(ctx context.Context) ([]PsychologicalType, error)
// func (q *TestRepository) ListQuestions(ctx context.Context) ([]Question, error)
// func (q *TestRepository) ListUsers(ctx context.Context) ([]User, error)

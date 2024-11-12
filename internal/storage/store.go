package storage

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	schemas "github.com/curtrika/UMetrika_server/internal/storage/schemas"
	"sync"

	"github.com/curtrika/UMetrika_server/internal/converter"
	"github.com/curtrika/UMetrika_server/internal/domain/models"
	storage "github.com/curtrika/UMetrika_server/internal/storage/sqlc_gen"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	_ "github.com/lib/pq"
)

// TODO: вынести в отдельный файл
type Storage struct {
	cvt converter.PsqlConverter
	db  *sql.DB
	*storage.Queries
}

func DatabaseInit(databaseURL string) (*Storage, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		return nil, err
	}

	//if err := db.Ping(); err != nil {
	//	return nil, err
	//}
	queries := storage.New(conn)

	return &Storage{
		db:      db,
		Queries: queries,
	}, nil
}

// TODO: CRUD вынести в отдельные модули
// SaveUser saves user to db.
func (s *Storage) SaveUser(ctx context.Context, email string, passHash []byte) (uuid.UUID, error) {
	const op = "storage.SaveUser"

	q := `insert into users (id, email, pass_hash)
	values ($1, $2, $3)
	returning id;`

	newUserId, err := uuid.NewUUID()
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	err = s.db.QueryRowContext(ctx, q, newUserId, email, passHash).Scan(&newUserId)
	if err != nil {
		// TODO: добавить кунг-фу ошибками (обработку на дубликаты)
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	return newUserId, nil
}

func (s *Storage) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return nil, nil
}

func (s *Storage) SaveApp(ctx context.Context, app models.App) (int32, error) {
	return 0, nil
}

func (s *Storage) GetAppById(ctx context.Context, appID int32) (*models.App, error) {
	const op = "storage.GetUserByEmail"

	q := `select json_build_object(
	    'id', id,
	    'name', name,
	    'secret', secret
	)
	from apps
	where id = $1;`

	var bs []byte
	if err := s.db.QueryRowContext(ctx, q, appID).Scan(&bs); err != nil {
		// TODO: добавить кунг-фу ошибками (sql no rows)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var schema schemas.AppSchema
	if err := json.Unmarshal(bs, &schema); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	appModel := s.cvt.AppToModel(schema)

	return &appModel, nil
}

type MockDatabase struct {
	mu         sync.Mutex
	usersTable map[string]models.User
	appTable   map[int32]models.App
	appId      int32
}

func NewMockDatabase() MockDatabase {
	return MockDatabase{
		usersTable: make(map[string]models.User),
		appTable:   make(map[int32]models.App),
	}
}

func (m *MockDatabase) SaveUser(ctx context.Context, email string, passHash []byte) (uuid.UUID, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	u := uuid.New()
	if _, ok := m.usersTable[email]; ok {
		return uuid.UUID{}, fmt.Errorf("err: user already exists")
	}
	m.usersTable[u.String()] = models.User{
		ID:       u,
		Email:    email,
		PassHash: passHash,
	}
	m.usersTable[email] = models.User{
		ID:       u,
		Email:    email,
		PassHash: passHash,
	}
	return u, nil
}

func (m *MockDatabase) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	u := m.usersTable[email]
	return &u, nil
}

func (m *MockDatabase) GetAppById(ctx context.Context, appID int32) (*models.App, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	app := m.appTable[appID]
	return &app, nil
}

func (m *MockDatabase) SaveApp(ctx context.Context, app models.App) (int32, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.appTable[m.appId] = app
	id := m.appId
	m.appId += 1
	return id, nil
}

func (m *MockDatabase) SaveAppById(ctx context.Context, id int32, app models.App) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.appTable[id] = app
}

func (m *MockDatabase) CreateAnswer(ctx context.Context, arg storage.CreateAnswerParams) (storage.Answer, error) {
	// Mock response to simulate a created answer
	return storage.Answer{
		ID:           1, // Mocked ID
		NextAnswerID: arg.NextAnswerID,
		Title:        arg.Title,
	}, nil
}

func (m *MockDatabase) CreateApp(ctx context.Context, arg storage.CreateAppParams) (storage.App, error) {
	// Mock response to simulate a created app
	return storage.App{
		ID:     1,          // Mocked ID
		Name:   arg.Name,   // Name from params
		Secret: arg.Secret, // Secret from params
	}, nil
}

func (m *MockDatabase) CreatePsychologicalPerformance(ctx context.Context, arg storage.CreatePsychologicalPerformanceParams) (storage.PsychologicalPerformance, error) {
	// Mock response to simulate a created psychological performance
	return storage.PsychologicalPerformance{
		ID:                  1,                       // Mocked ID
		OwnerID:             arg.OwnerID,             // Owner ID from params
		PsychologicalTestID: arg.PsychologicalTestID, // Test ID from params
		StartedAt:           arg.StartedAt,           // Started at timestamp
	}, nil
}

func (m *MockDatabase) CreatePsychologicalTest(ctx context.Context, arg storage.CreatePsychologicalTestParams) (storage.PsychologicalTest, error) {
	// Mock response to simulate a created psychological test
	return storage.PsychologicalTest{
		ID:              1,                   // Mocked ID
		FirstQuestionID: arg.FirstQuestionID, // First question ID from params
		TypeID:          arg.TypeID,          // Type ID from params
		OwnerID:         arg.OwnerID,         // Owner ID from params
		Title:           arg.Title,           // Title from params
	}, nil
}

func (m *MockDatabase) CreatePsychologicalType(ctx context.Context, title string) (storage.PsychologicalType, error) {
	// Mock response to simulate a created psychological type
	return storage.PsychologicalType{
		ID:    1,     // Mocked ID
		Title: title, // Title from params
	}, nil
}

func (m *MockDatabase) CreateQuestion(ctx context.Context, arg storage.CreateQuestionParams) (storage.Question, error) {
	// Mock response to simulate a created question
	return storage.Question{
		ID:             1,                  // Mocked ID
		NextQuestionID: arg.NextQuestionID, // Next question ID from params
		Number:         arg.Number,         // Question number from params
		FirstAnswerID:  arg.FirstAnswerID,  // First answer ID from params
		Title:          arg.Title,          // Title from params
	}, nil
}

//func (m *MockDatabase) CreateUser(ctx context.Context, arg storage.CreateUserParams) (storage.User, error) {
//	// Mock response to simulate a created user
//	return storage.User{
//		ID:       pgtype.UUID{}, // ID from params (could be UUID)
//		Email:    arg.Email,     // Email from params
//		PassHash: arg.PassHash,  // Pass hash from params
//	}, nil
//}

func (m *MockDatabase) GetAnswer(ctx context.Context, id int32) (storage.Answer, error) {
	// Mock response for fetching an answer
	return storage.Answer{
		ID:           id,                    // Use provided ID
		NextAnswerID: pgtype.Int4{Int32: 1}, // For simplicity, mock next_answer_id as nil
		Title:        "Mock Answer",         // Mock title
	}, nil
}

func (m *MockDatabase) GetApp(ctx context.Context, id int32) (storage.App, error) {
	// Mock response for fetching an app
	return storage.App{
		ID:     id,          // Use provided ID
		Name:   "Mock App",  // Mock name
		Secret: "SecretKey", // Mock secret
	}, nil
}

func (m *MockDatabase) GetPsychologicalPerformance(ctx context.Context, id int32) (storage.PsychologicalPerformance, error) {
	// Mock response for fetching psychological performance
	return storage.PsychologicalPerformance{
		ID:                  id,                    // Use provided ID
		OwnerID:             1,                     // Mocked owner ID
		PsychologicalTestID: pgtype.Int4{Int32: 1}, // Mocked test ID
		StartedAt:           pgtype.Timestamptz{},  // Current time
	}, nil
}

func (m *MockDatabase) GetPsychologicalTest(ctx context.Context, id int32) (storage.PsychologicalTest, error) {
	// Mock response for fetching a psychological test
	return storage.PsychologicalTest{
		ID:              id,            // Use provided ID
		FirstQuestionID: pgtype.Int4{}, // Mocked first question ID
		TypeID:          pgtype.Int4{}, // Mocked type ID
		OwnerID:         1,             // Mocked owner ID
		Title:           "Mock Test",   // Mocked title
	}, nil
}

func (m *MockDatabase) GetPsychologicalType(ctx context.Context, id int32) (storage.PsychologicalType, error) {
	// Mock response for fetching psychological type
	return storage.PsychologicalType{
		ID:    id,          // Use provided ID
		Title: "Mock Type", // Mocked title
	}, nil
}

func (m *MockDatabase) GetQuestion(ctx context.Context, id int32) (storage.Question, error) {
	// Mock response for fetching a question
	return storage.Question{
		ID:             id,              // Use provided ID
		Number:         1,               // Mock question number
		NextQuestionID: pgtype.Int4{},   // For simplicity, mock as nil
		FirstAnswerID:  pgtype.Int4{},   // For simplicity, mock as nil
		Title:          "Mock Question", // Mocked title
	}, nil
}

func (m *MockDatabase) GetUser(ctx context.Context, id pgtype.UUID) (storage.User, error) {
	// Mock response for fetching a user
	return storage.User{
		ID:       id,                       // Mocked UUID
		Email:    "user@example.com",       // Mocked email
		PassHash: []byte("hashedpassword"), // Mocked password hash
	}, nil
}

func (m *MockDatabase) ListAnswers(ctx context.Context) ([]storage.Answer, error) {
	// Mocked list of answers
	return []storage.Answer{
		{ID: 1, Title: "Answer 1"},
		{ID: 2, Title: "Answer 2"},
	}, nil
}

func (m *MockDatabase) ListApps(ctx context.Context) ([]storage.App, error) {
	// Mocked list of apps
	return []storage.App{
		{ID: 1, Name: "App 1", Secret: "Secret1"},
		{ID: 2, Name: "App 2", Secret: "Secret2"},
	}, nil
}

func (m *MockDatabase) ListPsychologicalPerformances(ctx context.Context) ([]storage.PsychologicalPerformance, error) {
	// Mocked list of psychological performances
	return []storage.PsychologicalPerformance{
		{ID: 1, OwnerID: 1, PsychologicalTestID: pgtype.Int4{}, StartedAt: pgtype.Timestamptz{}},
	}, nil
}

func (m *MockDatabase) ListPsychologicalTests(ctx context.Context) ([]storage.PsychologicalTest, error) {
	// Mocked list of psychological tests
	return []storage.PsychologicalTest{
		{ID: 1, FirstQuestionID: pgtype.Int4{}, TypeID: pgtype.Int4{}, OwnerID: 1, Title: "Test 1"},
	}, nil
}

func (m *MockDatabase) ListPsychologicalTypes(ctx context.Context) ([]storage.PsychologicalType, error) {
	// Mocked list of psychological types
	return []storage.PsychologicalType{
		{ID: 1, Title: "Type 1"},
	}, nil
}

func (m *MockDatabase) ListQuestions(ctx context.Context) ([]storage.Question, error) {
	// Mocked list of questions
	return []storage.Question{
		{ID: 1, Number: 1, Title: "Question 1"},
	}, nil
}

func (m *MockDatabase) ListUsers(ctx context.Context) ([]storage.User, error) {
	// Mocked list of users
	return []storage.User{
		{ID: pgtype.UUID{}, Email: "user1@example.com"},
		{ID: pgtype.UUID{}, Email: "user2@example.com"},
	}, nil
}

func (m *MockDatabase) WithTx(tx pgx.Tx) *storage.Queries {
	// Mocked implementation of WithTx
	return &storage.Queries{}
}

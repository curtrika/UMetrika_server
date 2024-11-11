package storage

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	storage "github.com/curtrika/UMetrika_server/internal/storage/schemas"
	"sync"

	"github.com/curtrika/UMetrika_server/internal/converter"
	"github.com/curtrika/UMetrika_server/internal/domain/models"
	"github.com/google/uuid"

	_ "github.com/lib/pq"
)

// TODO: вынести в отдельный файл
type Storage struct {
	cvt converter.PsqlConverter
	db  *sql.DB
}

func DatabaseInit(databaseURL string) (*Storage, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	//if err := db.Ping(); err != nil {
	//	return nil, err
	//}

	return &Storage{
		db: db,
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

func (s *Storage) GetAppById(ctx context.Context, appID uuid.UUID) (*models.App, error) {
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

	var schema storage.AppSchema
	if err := json.Unmarshal(bs, &schema); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	appModel := s.cvt.AppToModel(schema)

	return &appModel, nil
}

func (s *Storage) CreateUser(ctx context.Context, user models.User) (*models.User, error) {
	const op = "storage.CreateUser"

	params := []any{user.FirstName, user.MiddleName, user.LastName, user.Email, user.PassHash, user.RoleTitle, user.SchoolID, user.ClassesID}
	q := `INSERT INTO users (first_name, middle_name, last_name, email, role_title, school_id, classes_id)
		VALUES (:id, :first_name, :middle_name, :last_name, :email, :pass_hash, :role_title, :school_id, :classes_id)
		ON CONFLICT DO NOTHING;`

	var bs []byte
	if err := s.db.QueryRowContext(ctx, q, params).Scan(&bs); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var schema storage.UserSchema
	if err := json.Unmarshal(bs, &schema); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	userModel := s.cvt.UserToModel(schema)

	return &userModel, nil
}

func (s *Storage) ReadUser(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	const op = "storage.ReadUser"

	q := `SELECT * FROM users WHERE id = $1;`

	var bs []byte
	if err := s.db.QueryRowContext(ctx, q, userID).Scan(&bs); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var schema storage.UserSchema
	if err := json.Unmarshal(bs, &schema); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	userModel := s.cvt.UserToModel(schema)

	return &userModel, nil
}

func (s *Storage) UpdateUser(ctx context.Context, user models.User) (*models.User, error) {
	return nil, nil
}

func (s *Storage) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	return nil
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

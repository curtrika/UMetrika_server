package storage

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	"github.com/curtrika/UMetrika_server/internal/domain/models"
	"github.com/curtrika/UMetrika_server/internal/storage/schemas"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

// TODO: вынести в отдельный файл
type Storage struct {
	cvt Converter
	db  *sql.DB
}

//func DatabaseInit(databaseURL string) (*Storage, error) {
//	db, err := sql.Open("postgres", databaseURL)
//	if err != nil {
//		return nil, err
//	}
//	conn, err := pgx.Connect(context.Background(), databaseURL)
//	if err != nil {
//		return nil, err
//	}
//
//	//if err := db.Ping(); err != nil {
//	//	return nil, err
//	//}
//	queries := storage.New(conn)
//
//	return &Storage{
//		db: db,
//	}, nil
//}

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

	//appModel := s.cvt.AppToModel(schema)

	return nil, nil
}

func (s *Storage) GetTeacherDisciplinesAndClasses(ctx context.Context, teacherID uuid.UUID) ([]models.TeacherDiscipline, error) {
	const op = "postgres.GetTeacherDisciplinesAndClasses"

	path := "./queries/GetTeacherDisciplinesAndClasses.sql"
	q, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("%s: ошибка при чтении файла: %s: %w", op, path, err)
	}

	var bs []byte
	if err := s.db.QueryRowContext(ctx, string(q), teacherID).Scan(&bs); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var TeacherDisciplinesSchema []schemas.TeacherDisciplineSchema
	if err := json.Unmarshal(bs, &TeacherDisciplinesSchema); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	teacherDisciplines := s.cvt.TeacherDisciplinesToModel(TeacherDisciplinesSchema)

	return teacherDisciplines, nil
}

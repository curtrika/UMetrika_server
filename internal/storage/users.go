package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/curtrika/UMetrika_server/internal/domain/models"
	storage "github.com/curtrika/UMetrika_server/internal/storage/schemas"
)

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

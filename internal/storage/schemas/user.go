package schemas

import (
	"github.com/google/uuid"
	"time"
)

type UserSchema struct {
	ID        uuid.UUID `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	PassHash  []byte    `json:"pass_hash"`
	RoleID    uuid.UUID `json:"role_id"`
	SchoolID  uuid.UUID `json:"school_id"`
	ClassesID uuid.UUID `json:"classes_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

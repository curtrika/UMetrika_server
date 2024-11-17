package schemas

import (
	"github.com/google/uuid"
	"time"
)

type UserSchema struct {
	ID         uuid.UUID `json:"id"`
	FirstName  string    `json:"first_name"`
	MiddleName string    `json:"middle_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	Gender     bool      `json:"gender"`
	Role       string    `json:"role"`
	SchoolID   uuid.UUID `json:"school_id"`
	ClassesID  uuid.UUID `json:"classes_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}

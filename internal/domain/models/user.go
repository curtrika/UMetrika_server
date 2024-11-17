package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID         uuid.UUID
	FirstName  string
	MiddleName string
	LastName   string
	Email      string
	PassHash   []byte
	Gender     bool
	RoleID     uuid.UUID
	SchoolID   uuid.UUID
	ClassesID  uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

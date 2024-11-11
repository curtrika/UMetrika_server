package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id        uuid.UUID
	FullName  string
	Email     string
	Password  string
	PassHash  []byte
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

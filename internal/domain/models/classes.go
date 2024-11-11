package models

import (
	"github.com/google/uuid"
	"time"
)

type Classes struct {
	ID            uuid.UUID
	Grade         int
	Title         string
	MainTeacherID uuid.UUID
	ReleaseDate   time.Time
}

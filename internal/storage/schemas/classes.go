package storage

import (
	"github.com/google/uuid"
	"time"
)

type ClassesSchema struct {
	ID            uuid.UUID `json:"id"`
	Grade         int       `json:"grade"`
	Title         string    `json:"title"`
	MainTeacherID uuid.UUID `json:"main_teacher_id"`
	ReleaseDate   time.Time `json:"release_date"`
}

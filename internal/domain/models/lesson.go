package models

import "github.com/google/uuid"

type Lesson struct {
	ID        uuid.UUID
	ThemeID   uuid.UUID
	GroupID   uuid.UUID
	TeacherID uuid.UUID
}

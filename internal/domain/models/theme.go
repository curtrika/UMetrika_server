package models

import "github.com/google/uuid"

type Theme struct {
	ID           uuid.UUID
	Title        string
	DisciplineID uuid.UUID
}

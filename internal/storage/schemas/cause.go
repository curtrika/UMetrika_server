package storage

import "github.com/google/uuid"

type CauseSchema struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	ProblemID string    `json:"problem_id"`
}

package storage

import "github.com/google/uuid"

type SolutionSchema struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	ProblemID uuid.UUID `json:"problem_id"`
}

package schemas

import "github.com/google/uuid"

type Solution struct {
	ID        uuid.UUID
	Title     string
	ProblemID uuid.UUID
}

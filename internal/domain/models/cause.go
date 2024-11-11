package schemas

import "github.com/google/uuid"

type Cause struct {
	ID        uuid.UUID
	Title     string
	ProblemID string
}

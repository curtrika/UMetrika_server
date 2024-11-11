package schemas

import "github.com/google/uuid"

type ProblemSchema struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
}

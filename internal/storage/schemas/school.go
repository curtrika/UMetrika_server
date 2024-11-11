package schemas

import "github.com/google/uuid"

type SchoolSchema struct {
	ID        uuid.UUID `json:"id"`
	LargeName string    `json:"large_name"`
}

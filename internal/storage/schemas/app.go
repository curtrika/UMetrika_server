package schemas

import "github.com/google/uuid"

type AppSchema struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Secret string    `json:"secret"`
}

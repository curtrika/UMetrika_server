package storage

import "github.com/google/uuid"

type RoleSchema struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
}

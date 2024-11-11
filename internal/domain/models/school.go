package models

import "github.com/google/uuid"

type School struct {
	ID        uuid.UUID
	LargeName string
}

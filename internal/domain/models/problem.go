package schemas

import "github.com/google/uuid"

type Problem struct {
	ID    uuid.UUID
	Title string
}
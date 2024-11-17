package models

import "github.com/google/uuid"

type Classes struct {
	ID            uuid.UUID
	Title         string
	MainTeacherID uuid.UUID
	Students      []User
}

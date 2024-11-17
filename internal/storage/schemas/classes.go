package schemas

import "github.com/google/uuid"

type ClassesSchema struct {
	ID            uuid.UUID    `json:"id"`
	Title         string       `json:"title"`
	MainTeacherID uuid.UUID    `json:"main_teacher_id"`
	Students      []UserSchema `json:"students"`
}

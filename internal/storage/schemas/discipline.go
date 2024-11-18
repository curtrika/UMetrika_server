package schemas

import "github.com/google/uuid"

type DisciplineSchema struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type TeacherDisciplineSchema struct {
	ID      uuid.UUID       `json:"id"`
	Title   uuid.UUID       `json:"title"`
	Classes []ClassesSchema `json:"classes"`
}

package models

import "github.com/google/uuid"

type Discipline struct {
	ID   uuid.UUID
	Name string
}

type TeacherDiscipline struct {
	ID      uuid.UUID
	Title   uuid.UUID
	Classes []Classes
}

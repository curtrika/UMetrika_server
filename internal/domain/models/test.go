package models

import (
	"time"

	"github.com/google/uuid"
)

type EducationOwner struct {
	OwnerID   uuid.UUID
	OwnerName string
	PassHash  []byte
	Email     string
	CreatedAt time.Time
}

type EducationTest struct {
	TestID      uuid.UUID
	OwnerID     uuid.UUID
	TestName    string
	Description string
	TestType    string
	CreatedAt   time.Time
}

// deprecate

type PsychologicalPerformance struct {
	ID                  int32
	OwnerID             int32
	PsychologicalTestID int
	StartedAt           time.Time
}

type PsychologicalTest struct {
	ID              int32
	FirstQuestionID int
	TypeID          int
	OwnerID         int32
	Title           string
}

type Answer struct {
	ID           int32
	NextAnswerID int
	Title        string
}

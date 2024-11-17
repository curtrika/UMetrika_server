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

type EducationTestFull struct {
	TestID      uuid.UUID
	OwnerID     uuid.UUID
	TestName    string
	Description string
	TestType    string
	CreatedAt   time.Time
	Questions   []*EducationQuestionFull
}

type EducationQuestion struct {
	QuestionID    uuid.UUID
	TestID        uuid.UUID
	QuestionText  string
	QuestionType  string
	QuestionOrder int32
	CreatedAt     time.Time
}

type EducationQuestionFull struct {
	QuestionID    uuid.UUID
	TestID        uuid.UUID
	QuestionText  string
	QuestionType  string
	QuestionOrder int32
	CreatedAt     time.Time
	Answers       []EducationAnswer
}

type QuestionAnswer struct {
	TestID    uuid.UUID
	Questions EducationQuestion
	Answers   []EducationAnswer
}

type EducationAnswer struct {
	AnswerID    uuid.UUID
	QuestionID  uuid.UUID
	AnswerText  string
	AnswerOrder int32
	ScoreValue  int
	CreatedAt   time.Time
	// IsCorrect *bool
}

// type EducationTestWithQuestions struct {
// 	EducationTest
// 	Questions []EducationQuestion
// 	// IsGraded  bool
// }

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

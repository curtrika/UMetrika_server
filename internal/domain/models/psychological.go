package models

import "time"

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

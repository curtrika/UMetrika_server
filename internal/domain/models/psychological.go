package models

import "time"

type PsychologicalPerformance struct {
	ID                  int32     `json:"id"`
	OwnerID             int32     `json:"owner_id"`
	PsychologicalTestID int       `json:"psychological_test_id"`
	StartedAt           time.Time `json:"started_at"`
}

package postgres

import (
	"time"

	"github.com/curtrika/UMetrika_server/internal/domain/models"
	"github.com/curtrika/UMetrika_server/internal/repository/postgres/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

// goverter:converter
// goverter:extend Int4ToInt
// goverter:extend IntToInt4
// goverter:extend TimestamptzToTime
// goverter:extend TimeToTimestamptz
type Converter interface {
	PsychologicalPerfomanceDBToModel(dbModel postgres.PsychologicalPerformance) models.PsychologicalPerformance
	PsychologicalPerfomanceModelToDB(model models.PsychologicalPerformance) postgres.PsychologicalPerformance

	PsychologicalTestDBToModel(dbModel postgres.PsychologicalTest) models.PsychologicalTest
	PsychologicalTestModelToDB(model models.PsychologicalTest) postgres.PsychologicalTest

	PsychologicalTestsDBToModel(dbModel []postgres.PsychologicalTest) []models.PsychologicalTest
	PsychologicalTestsModelToDB(model []models.PsychologicalTest) []postgres.PsychologicalTest

	AnswerDBToModel(dbModel postgres.Answer) models.Answer
	AnswerModelToDB(model models.Answer) postgres.Answer
}

func Int4ToInt(val pgtype.Int4) int {
	return int(val.Int32)
}

func TimestamptzToTime(val pgtype.Timestamptz) time.Time {
	return val.Time
}

func IntToInt4(val int) pgtype.Int4 {
	ret := pgtype.Int4{}
	ret.Scan(&val)
	return ret
}

func TimeToTimestamptz(val time.Time) pgtype.Timestamptz {
	ret := pgtype.Timestamptz{}
	ret.Scan(&val)
	return ret
}

package storage

import (
	"time"

	"github.com/curtrika/UMetrika_server/internal/domain/models"
	"github.com/jackc/pgx/v5/pgtype"
)

// goverter:converter
// goverter:extend Int4ToInt
// goverter:extend IntToInt4
// goverter:extend TimestamptzToTime
// goverter:extend TimeToTimestamptz
type Converter interface {
	DBToModel(dbModel *PsychologicalPerformance) (*models.PsychologicalPerformance, error)
	ModelToDB(model *models.PsychologicalPerformance) (*PsychologicalPerformance, error)
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

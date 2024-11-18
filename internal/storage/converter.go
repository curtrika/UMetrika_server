package storage

import (
	"github.com/curtrika/UMetrika_server/internal/domain/models"
	"github.com/curtrika/UMetrika_server/internal/storage/schemas"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

// GrpcConverter is an interface for converting between gRPC and service layer representations.
// goverter:converter
// goverter:output:file ./generated/converter.gen.go
// goverter:output:package :storage
// goverter:useZeroValueOnPointerInconsistency
// goverter:ignoreUnexported
// goverter:matchIgnoreCase
// goverter:extend UUIDtoUUID
//
//go:generate goverter gen ./
type Converter interface {
	AppToModel(request schemas.AppSchema) (response models.App)
	TeacherDisciplinesToModel(request []schemas.TeacherDisciplineSchema) (response []models.TeacherDiscipline)
}

func UUIDtoUUID(source uuid.UUID) uuid.UUID {
	return source
}

func UUIDPostgresToGoogle(v pgtype.UUID) uuid.UUID {
	return v.Bytes
}

func UUIDGoogleToPostgres(v uuid.UUID) pgtype.UUID {
	var ret pgtype.UUID
	ret.Scan(v.String())
	return ret
}

func Int4ToInt(val pgtype.Int4) int {
	return int(val.Int32)
}

func TimestamptzToTime(val pgtype.Timestamptz) time.Time {
	return val.Time
}

func TimestampToTime(val pgtype.Timestamp) time.Time {
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

func TimeToTimestamp(val time.Time) pgtype.Timestamp {
	ret := pgtype.Timestamp{}
	ret.Scan(&val)
	return ret
}

func PostgresTextToString(t pgtype.Text) string {
	return t.String
}

func StringToPostgresText(s string) pgtype.Text {
	ret := pgtype.Text{}
	ret.Scan(&s)
	return ret
}

func NumericToInt(val pgtype.Numeric) int {
	return int(val.Int.Int64())
}

func IntToNumeric(val int) pgtype.Numeric {
	ret := pgtype.Numeric{}
	ret.Scan(&val)
	return ret
}

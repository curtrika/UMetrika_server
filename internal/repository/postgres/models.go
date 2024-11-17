package postgres

import (
	"time"

	"github.com/curtrika/UMetrika_server/internal/domain/models"
	postgres "github.com/curtrika/UMetrika_server/internal/repository/postgres/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// goverter:converter
// goverter:extend Int4ToInt
// goverter:extend IntToInt4
// goverter:extend TimestamptzToTime
// goverter:extend TimeToTimestamptz
// goverter:extend UUIDPostgresToGoogle
// goverter:extend UUIDGoogleToPostgres
// goverter:extend TimestampToTime
// goverter:extend TimeToTimestamp
// goverter:extend PostgresTextToString
// goverter:extend StringToPostgresText
// goverter:extend NumericToInt
// goverter:extend IntToNumeric
type Converter interface {
	OwnerDBToModel(dbModel postgres.EducationOwner) models.EducationOwner
	OwnerModelToDb(Model models.EducationOwner) postgres.EducationOwner
	TestDBToModel(dbModel postgres.EducationTest) models.EducationTest
	TestModelToDB(model models.EducationTest) postgres.EducationTest
	TestsDBToModel(dbModel []postgres.EducationTest) []models.EducationTest
	TestsModelToDB(model []models.EducationTest) []postgres.EducationTest

	QuestionModelToDB(model models.EducationQuestion) postgres.EducationQuestion
	QuestionDBToModel(dbModel postgres.EducationQuestion) models.EducationQuestion

	QuestionsModelToDB(model []models.EducationQuestion) []postgres.EducationQuestion
	QuestionsDBToModel(dbModel []postgres.EducationQuestion) []models.EducationQuestion

	AnswerModelToDB(model models.EducationAnswer) postgres.EducationAnswer
	AnswerDBToModel(dbModel postgres.EducationAnswer) models.EducationAnswer

	AnswersModelToDB(model []models.EducationAnswer) []postgres.EducationAnswer
	AnswersDBToModel(dbModel []postgres.EducationAnswer) []models.EducationAnswer

	TeacherDisciplineDBToModel(dbModel []postgres.TeacherDiscipline) []models.TeacherDiscipline

	// PsychologicalPerfomanceDBToModel(dbModel postgres.PsychologicalPerformance) models.PsychologicalPerformance
	// PsychologicalPerfomanceModelToDB(model models.PsychologicalPerformance) postgres.PsychologicalPerformance
	//
	// PsychologicalTestDBToModel(dbModel postgres.PsychologicalTest) models.PsychologicalTest
	// PsychologicalTestModelToDB(model models.PsychologicalTest) postgres.PsychologicalTest
	//
	// PsychologicalTestsDBToModel(dbModel []postgres.PsychologicalTest) []models.PsychologicalTest
	// PsychologicalTestsModelToDB(model []models.PsychologicalTest) []postgres.PsychologicalTest
	//
	// AnswerDBToModel(dbModel postgres.Answer) models.Answer
	// AnswerModelToDB(model models.Answer) postgres.Answer
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

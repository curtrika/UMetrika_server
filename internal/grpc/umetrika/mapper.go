package umetrikagrpc

import (
	"time"

	"github.com/curtrika/UMetrika_server/internal/domain/models"
	"github.com/curtrika/UMetrika_server/pkg/proto/umetrika/v1"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// goverter:converter
// goverter:output:package :converter
// goverter:useZeroValueOnPointerInconsistency
// goverter:ignoreUnexported
// goverter:matchIgnoreCase
// goverter:extend StringToUUID
// goverter:extend UUIDToString
// goverter:extend UUIDToString
// goverter:extend TimestampToTime
// goverter:extend TimeToTimeStamp
type Converter interface {
	// goverter:ignore OwnerID PassHash CreatedAt
	OwnerProtoToModel(proto *v1.OwnerPost) (*models.EducationOwner, error)
	// goverter:ignore state sizeCache unknownFields
	// goverter:map OwnerID OwnerId
	OwnerModelToProto(m *models.EducationOwner) *v1.OwnerResult

	// goverter:ignore TestID CreatedAt
	// goverter:map OwnerId OwnerID
	TestProtoToModel(proto *v1.TestPost) (*models.EducationTest, error)

	// goverter:ignore state sizeCache unknownFields
	// goverter:map TestID TestId
	// goverter:map OwnerID OwnerId
	TestModelToProto(model *models.EducationTest) (*v1.TestResult, error)

	// goverter:ignore CreatedAt TestID QuestionID
	QuestionDTOProtoToModel(q *v1.QuestionPostDTO) (*models.EducationQuestion, error)

	// goverter:ignore CreatedAt AnswerID ScoreValue QuestionID
	AnswerDTOProtoToModel(q *v1.AnswerPostDTO) (*models.EducationAnswer, error)

	QuestionDTOsProtoToModel(q []*v1.QuestionPostDTO) ([]*models.EducationQuestion, error)

	AnswerDTOsProtoToModel(q []*v1.AnswerPostDTO) ([]*models.EducationAnswer, error)
	// TestModelToProto(proto *v1.TestPost) (*models.EducationTestWithQuestions, error)

	TeacherDisciplineToProto(model models.TeacherDiscipline) (*v1.TeacherDiscipline, error)
}

func StringToUUID(s string) (uuid.UUID, error) {
	return uuid.Parse(s)
}

func UUIDToString(u uuid.UUID) string {
	return u.String()
}

func TimestampToTime(t *timestamppb.Timestamp) time.Time {
	return t.AsTime()
}

func TimeToTimeStamp(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}

package umetrikagrpc

import (
	"time"

	"github.com/curtrika/UMetrika_server/internal/domain/models"
	"github.com/curtrika/UMetrika_server/pkg/proto/umetrika/v1"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// goverter:converter
// goverter:extend StringToUUID
// goverter:extend UUIDToString
// goverter:extend UUIDToString
// goverter:extend TimestampToTime
// goverter:extend TimeToTimeStamp
type Converter interface {
	// goverter:ignore OwnerID PassHash CreatedAt
	OwnerProtoToModel(proto *v1.OwnerPost) (*models.EducationOwner, error)
	// goverter:ignore state sizeCache unknownFields
	OwnerModelToProto(m *models.EducationOwner) *v1.OwnerResult
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

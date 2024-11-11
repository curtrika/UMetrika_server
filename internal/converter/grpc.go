package converter

import (
	"github.com/curtrika/UMetrika_server/internal/domain/models"
	adminpanelv1 "github.com/curtrika/UMetrika_server/pkg/proto/admin_panel/v1"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

// GrpcConverter is an interface for converting between gRPC and service layer representations.
// goverter:converter
// goverter:output:file ./converter/converter.gen.go
// goverter:output:package :converter
// goverter:useZeroValueOnPointerInconsistency
// goverter:ignoreUnexported
// goverter:matchIgnoreCase
// goverter:extend StringToUUID
// goverter:extend UUIDToString
// goverter:extend TimeToTimestamp
//
//go:generate goverter gen ./
type GRPCConverter interface {
	// goverter:ignore PassHash
	UserToModel(request *adminpanelv1.User) (response *models.User)
	// goverter:ignore Password PassHash
	ModelToUser(request *models.User) (response *adminpanelv1.User)
}

func StringToUUID(source string) uuid.UUID {
	id, err := uuid.Parse(source)
	if err != nil {
		return uuid.Nil
	}
	return id
}

func UUIDToString(source uuid.UUID) string {
	if source == uuid.Nil {
		return ""
	}
	return source.String()
}

func TimeToTimestamp(t time.Time) *timestamppb.Timestamp {
	timestamp := timestamppb.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}
	return &timestamp
}

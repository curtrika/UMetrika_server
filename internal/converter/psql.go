package converter

import (
	"github.com/curtrika/UMetrika_server/internal/domain/models"
	"github.com/curtrika/UMetrika_server/internal/storage/schemas"
	"github.com/google/uuid"
)

// GrpcConverter is an interface for converting between gRPC and service layer representations.
// goverter:converter
// goverter:output:file ./converter/converter.gen.go
// goverter:output:package :converter
// goverter:useZeroValueOnPointerInconsistency
// goverter:ignoreUnexported
// goverter:matchIgnoreCase
// goverter:extend UUIDToUUID
//
//go:generate goverter gen ./
type PsqlConverter interface {
	UserToModel(request storage.UserSchema) (response models.User)

	AppToModel(request storage.AppSchema) (response models.App)
}

func UUIDToUUID(source uuid.UUID) uuid.UUID {
	return source
}

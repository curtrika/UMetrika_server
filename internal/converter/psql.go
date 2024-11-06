package converter

import (
	"github.com/curtrika/UMetrika_server/internal/domain/models"
	"github.com/curtrika/UMetrika_server/internal/storage/schemas"
)

// GrpcConverter is an interface for converting between gRPC and service layer representations.
// goverter:converter
// goverter:output:file ./converter/converter.gen.go
// goverter:output:package :converter
// goverter:useZeroValueOnPointerInconsistency
// goverter:ignoreUnexported
// goverter:matchIgnoreCase
//
//go:generate goverter gen ./
type PsqlConverter interface {
	AppToModel(request schemas.AppSchema) (response models.App)
}

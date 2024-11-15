package storage

import storage "github.com/curtrika/UMetrika_server/internal/storage/sqlc_gen"

type ConverterWrapper struct {
	*storage.Queries
	converter storage.Converter
}

func New(dbtx storage.DBTX, converter storage.Converter) ConverterWrapper {
	return ConverterWrapper{
		Queries:   storage.New(dbtx),
		converter: converter,
	}
}

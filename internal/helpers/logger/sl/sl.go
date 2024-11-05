package sl

import "log/slog"

// The function for adding an error to the log
// Usage example:
//
//	if err != nil {
//		a.log.Error("failed to get user", sl.Err(err))
//	}
func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}

package util

import (
	"io"

	"github.com/go-kit/kit/log"
)

func NewLogger(writer io.Writer) log.Logger {
	var logger log.Logger
	logger = log.NewLogfmtLogger(writer)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)
	return logger
}

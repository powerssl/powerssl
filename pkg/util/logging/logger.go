package logging

import (
	"os"

	"github.com/go-kit/kit/log"
)

func NewLogger() log.Logger {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	return logger
}
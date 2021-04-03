package common // import "powerssl.dev/common"

import (
	"io"

	"github.com/go-kit/kit/log"
	kitlogzap "github.com/go-kit/kit/log/zap"
	"go.uber.org/zap"
)

func NewLogger(writer io.Writer) log.Logger {
	var logger log.Logger
	logger = log.NewLogfmtLogger(writer)
	logger = withKeyVals(logger)
	return logger
}

func NewZapAndKitLogger() (_ *zap.Logger, _ log.Logger) {
	var zapLogger *zap.Logger
	var err error
	if zapLogger, err = zap.NewDevelopment(); err != nil {
		panic(err)
	}
	logger := kitlogzap.NewZapSugarLogger(zapLogger, zap.DebugLevel)
	logger = withKeyVals(logger)
	return zapLogger, logger
}

func withKeyVals(logger log.Logger) log.Logger {
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)
	return logger
}

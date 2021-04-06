package log // import "powerssl.dev/common/log"

import (
	kitlog "github.com/go-kit/kit/log"
	kitzap "github.com/go-kit/kit/log/zap"
	"go.uber.org/zap"
)

type Logger = *zap.SugaredLogger

func NewLogger(production bool, options ...zap.Option) (_ *zap.SugaredLogger, err error) {
	var config zap.Config
	if production {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
	}
	config.OutputPaths = []string{"stdout"}
	var logger *zap.Logger
	if logger, err = config.Build(options...); err != nil {
		return nil, err
	}
	zap.ReplaceGlobals(logger)
	return logger.Sugar(), nil
}

func KitLogger(logger Logger) kitlog.Logger {
	return kitzap.NewZapSugarLogger(logger.Desugar(), zap.DebugLevel)
}

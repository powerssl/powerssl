package log // import "powerssl.dev/common/log"

import (
	kitlog "github.com/go-kit/kit/log"
	kitzap "github.com/go-kit/kit/log/zap"
	"go.uber.org/zap"
)

func NewLogger(cfg Config) (_ *zap.SugaredLogger, err error) {
	config := zapConfig(cfg)
	var logger *zap.Logger
	if logger, err = config.Build(cfg.Options...); err != nil {
		return nil, err
	}
	zap.ReplaceGlobals(logger)
	return logger.Sugar(), nil
}

func KitLogger(logger *zap.SugaredLogger) kitlog.Logger {
	return kitzap.NewZapSugarLogger(logger.Desugar(), zap.DebugLevel)
}

func zapConfig(cfg Config) (config zap.Config) {
	if cfg.Production() {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
	}
	config.OutputPaths = []string{"stdout"}
	return config
}

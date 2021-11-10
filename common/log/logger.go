package log // import "powerssl.dev/common/log"

import (
	"go.uber.org/zap"
)

func New(cfg Config) (_ *zap.SugaredLogger, err error) {
	config := zapConfig(cfg)
	var logger *zap.Logger
	if logger, err = config.Build(cfg.Options...); err != nil {
		return nil, err
	}
	zap.ReplaceGlobals(logger)
	return logger.Sugar(), nil
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

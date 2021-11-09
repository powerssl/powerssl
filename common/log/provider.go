package log // import "powerssl.dev/common/log"

import (
	"github.com/google/wire"
	"go.uber.org/zap"
)

var Provider = wire.NewSet(
	Provide,
)

func Provide(cfg Config) (*zap.SugaredLogger, func(), error) {
	logger, err := NewLogger(cfg)
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		if err = logger.Sync(); err != nil {
			logger.Error(err)
		}
	}
	return logger, cleanup, nil
}

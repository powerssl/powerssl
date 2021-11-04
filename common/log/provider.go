package log // import "powerssl.dev/common/log"

import (
	"github.com/google/wire"
	"go.uber.org/zap"
)

var Provider = wire.NewSet(
	ProvideLogger,
)

func ProvideLogger() (*zap.SugaredLogger, func(), error) {
	logger, err := NewLogger(false)
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

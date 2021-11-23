package log // import "powerssl.dev/common/log"

import (
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	Provide,
)

func Provide(cfg Config) (Logger, func(), error) {
	logger, err := New(cfg)
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

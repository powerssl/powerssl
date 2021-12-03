package log // import "powerssl.dev/common/log"

import (
	"errors"
	"fmt"
	"os"
	"syscall"

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
		if err = logger.Sync(); err != nil && !errors.Is(err, syscall.ENOTTY) {
			err = fmt.Errorf("failed syncing logger: %w", err)
			_, _ = fmt.Fprintln(os.Stdout, err)
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
	}
	return logger, cleanup, nil
}

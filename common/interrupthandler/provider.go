package interrupthandler // import "powerssl.dev/common/interrupthandler"

import (
	"context"

	"github.com/google/wire"

	"powerssl.dev/common/log"
)

var Provider = wire.NewSet(
	Provide,
)

type F func() error

func Provide(ctx context.Context, logger log.Logger) F {
	logger = logger.With("component", "interruptHandler")
	return func() error {
		return InterruptHandler(ctx, logger)
	}
}

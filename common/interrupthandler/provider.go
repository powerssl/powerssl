package interrupthandler // import "powerssl.dev/common/interrupthandler"

import (
	"context"

	"github.com/google/wire"
	"go.uber.org/zap"
)

var Provider = wire.NewSet(
	Provide,
)

type F func() error

func Provide(ctx context.Context, logger *zap.SugaredLogger) F {
	logger = logger.With("component", "interruptHandler")
	return func() error {
		return InterruptHandler(ctx, logger)
	}
}

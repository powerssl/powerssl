package interrupthandler // import "powerssl.dev/common/interrupthandler"

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

type InterruptError struct {
	os.Signal
}

func (interruptError InterruptError) Error() string {
	return interruptError.String()
}

func InterruptHandler(ctx context.Context, logger *zap.SugaredLogger) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	select {
	case sig := <-c:
		logger.Infow("interrupt signal received", "signal", sig)
		return InterruptError{Signal: sig}
	case <-ctx.Done():
		logger.Error(ctx.Err())
		return ctx.Err()
	}
}

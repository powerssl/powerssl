package common

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
)

type InterruptError struct {
	os.Signal
}

func (interruptError InterruptError) Error() string {
	return interruptError.String()
}

func InterruptHandler(ctx context.Context, logger log.Logger) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	select {
	case sig := <-c:
		logger.Log("signal", sig)
		return InterruptError{Signal: sig}
	case <-ctx.Done():
		logger.Log("err", ctx.Err())
		return ctx.Err()
	}
}

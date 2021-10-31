package common // import "powerssl.dev/common"

import (
	"context"
	"fmt"
	"io"

	"go.uber.org/zap"
)

type CloserWithCtx interface {
	Close(ctx context.Context) error
}

func ErrWrapCloser(closer io.Closer, wErr *error) {
	if err := closer.Close(); err != nil && *wErr != nil {
		*wErr = fmt.Errorf("%s: %w", err, *wErr)
	} else if err != nil {
		*wErr = err
	}
}

func ErrWrapCloserWithCtx(closer CloserWithCtx, ctx context.Context, wErr *error) {
	if err := closer.Close(ctx); err != nil && *wErr != nil {
		*wErr = fmt.Errorf("%s: %w", err, *wErr)
	} else if err != nil {
		*wErr = err
	}
}

func ErrWrapSync(logger *zap.SugaredLogger, wErr *error) {
	if err := logger.Sync(); err != nil && *wErr != nil {
		*wErr = fmt.Errorf("%s: %w", err, *wErr)
	} else if err != nil {
		*wErr = err
	}
}

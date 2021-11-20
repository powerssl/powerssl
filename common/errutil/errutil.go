package errutil // import "powerssl.dev/common/errutil"

import (
	"fmt"
	"io"

	"go.uber.org/zap"
)

func ErrWrapCloser(closer io.Closer, wErr *error) {
	if err := closer.Close(); err != nil && *wErr != nil {
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

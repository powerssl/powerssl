package middleware // import "powerssl.dev/sdk/integration/middleware"

import (
	apiv1 "powerssl.dev/api/controller/v1"
	"powerssl.dev/common/log"
)

type loggingMiddleware struct {
	next   Handler
	logger log.Logger
}

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Handler) Handler {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

func (h *loggingMiddleware) Handle(activity *apiv1.Activity) {
	h.logger.Infow("Received activity", "activity", activity.GetToken(), "name", activity.GetName())
	h.next.Handle(activity)
}

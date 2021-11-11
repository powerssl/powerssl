package middleware // import "powerssl.dev/sdk/integration/middleware"

import (
	"go.uber.org/zap"

	apiv1 "powerssl.dev/api/controller/v1"
)

type loggingMiddleware struct {
	next   Handler
	logger *zap.SugaredLogger
}

func LoggingMiddleware(logger *zap.SugaredLogger) Middleware {
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

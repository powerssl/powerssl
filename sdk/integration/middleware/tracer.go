package middleware // import "powerssl.dev/sdk/integration/middleware"

import (
	apiv1 "powerssl.dev/api/controller/v1"
)

type tracingMiddleware struct {
	next Handler
}

func TracingMiddleware() Middleware {
	return func(next Handler) Handler {
		return &tracingMiddleware{
			next: next,
		}
	}
}

func (h *tracingMiddleware) Handle(activity *apiv1.Activity) {
	h.next.Handle(activity)
}

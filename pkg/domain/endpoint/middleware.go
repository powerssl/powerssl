package endpoint

import (
	"context"
	"time"

	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

func LoggingMiddleware(logger log.Logger) kitendpoint.Middleware {
	return func(next kitendpoint.Endpoint) kitendpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {

			defer func(begin time.Time) {
				logger.Log("transport_error", err, "took", time.Since(begin))
			}(time.Now())
			return next(ctx, request)

		}
	}
}

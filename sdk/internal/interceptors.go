package internal

import (
	"context"

	"google.golang.org/grpc"

	"powerssl.dev/common/log"
	"powerssl.dev/common/telemetry"
)

func AuthInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func LoggerInterceptor(logger log.Logger) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		logger.Infow("LOGGER", method, req, reply)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func TelemetryInterceptor(telemetry *telemetry.Telemeter) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		_ = telemetry.Tracer
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

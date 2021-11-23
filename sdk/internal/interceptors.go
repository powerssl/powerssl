package internal

import (
	"context"

	stdopentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	"powerssl.dev/common/log"
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

func TracerInterceptor(tracer stdopentracing.Tracer) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		_ = tracer
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

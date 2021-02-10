package apiserver

import (
	"context"

	"powerssl.dev/powerssl/pkg/apiserver/client"
)

var clientValue = struct{}{}

func GetClient(ctx context.Context) *client.GRPCClient {
	return ctx.Value(clientValue).(*client.GRPCClient)
}

func SetClient(ctx context.Context, client *client.GRPCClient) context.Context {
	return context.WithValue(ctx, clientValue, client)
}

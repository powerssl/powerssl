package apiserver

import (
	"context"

	"powerssl.dev/sdk/apiserver"
)

var clientValue = struct{}{}

func GetClient(ctx context.Context) *apiserver.Client {
	return ctx.Value(clientValue).(*apiserver.Client)
}

func SetClient(ctx context.Context, client *apiserver.Client) context.Context {
	return context.WithValue(ctx, clientValue, client)
}

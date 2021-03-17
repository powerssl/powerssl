package apiserver

import (
	"context"

	"powerssl.dev/backend/ctxkey"
	"powerssl.dev/sdk/apiserver"
)

var clientValue = ctxkey.New("dev.powerssl.backend.apiserver")

func GetClient(ctx context.Context) *apiserver.Client {
	return ctx.Value(clientValue).(*apiserver.Client)
}

func SetClient(ctx context.Context, client *apiserver.Client) context.Context {
	return context.WithValue(ctx, clientValue, client)
}

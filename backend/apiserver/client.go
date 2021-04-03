package apiserver // import "powerssl.dev/backend/apiserver"

import (
	"context"

	"powerssl.dev/sdk/apiserver"

	"powerssl.dev/backend/ctxkey"
)

var clientValue = ctxkey.New("dev.powerssl.backend.apiserver")

func GetClient(ctx context.Context) *apiserver.Client {
	return ctx.Value(clientValue).(*apiserver.Client)
}

func SetClient(ctx context.Context, client *apiserver.Client) context.Context {
	return context.WithValue(ctx, clientValue, client)
}

package sdk

import (
	"context"

	"powerssl.dev/sdk/apiserver"
)

type ctxKey struct {
	name string
}

var apiClient = ctxKey{name: "dev.powerssl.backend.apiserver"}

func GetAPIClient(ctx context.Context) *apiserver.Client {
	return ctx.Value(apiClient).(*apiserver.Client)
}

func SetAPIClient(ctx context.Context, client *apiserver.Client) context.Context {
	return context.WithValue(ctx, apiClient, client)
}

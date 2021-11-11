package context

import (
	"context"

	"powerssl.dev/sdk/apiserver"
)

var apiClient = New("dev.powerssl.backend.apiserver")

func GetAPIClient(ctx context.Context) *apiserver.Client {
	return ctx.Value(apiClient).(*apiserver.Client)
}

func SetAPIClient(ctx context.Context, client *apiserver.Client) context.Context {
	return context.WithValue(ctx, apiClient, client)
}

package context

import (
	"context"

	"powerssl.dev/backend/vault"
)

var vaultClient = New("dev.powerssl.backend.vault")

func GetVaultClient(ctx context.Context) *vault.Client {
	return ctx.Value(vaultClient).(*vault.Client)
}

func SetVaultClient(ctx context.Context, client *vault.Client) context.Context {
	return context.WithValue(ctx, vaultClient, client)
}

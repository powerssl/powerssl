package activity

import (
	"context"
	"strings"

	temporalactivity "go.temporal.io/sdk/activity"

	"powerssl.dev/powerssl/internal/pkg/vault"
)

type CreateVaultTransitKeyParams struct {
	Name string
}

func (p *CreateVaultTransitKeyParams) ToKeyVals() []interface{} {
	return []interface{}{
		"Name", p.Name,
	}
}

func CreateVaultTransitKey(ctx context.Context, params *CreateVaultTransitKeyParams) (err error) {
	logger := temporalactivity.GetLogger(ctx)
	logger.Info("CreateVaultTransitKey", params.ToKeyVals()...)

	name := strings.ReplaceAll(params.Name, "/", "-")
	logger.Info("Creating Vault Transit Key", "Name", name)
	return vault.GetClient(ctx).CreateTransitKey(ctx, name)
}

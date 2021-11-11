package activity

import (
	"context"
	"strings"

	temporalactivity "go.temporal.io/sdk/activity"

	context2 "powerssl.dev/backend/context"
	"powerssl.dev/workflow/activity"
)

func CreateVaultTransitKey(ctx context.Context, params *activity.CreateVaultTransitKeyParams) (err error) {
	logger := temporalactivity.GetLogger(ctx)
	logger.Info("CreateVaultTransitKey", params.ToKeyVals()...)

	name := strings.ReplaceAll(params.Name, "/", "-")
	logger.Info("Creating Vault Transit Key", "Name", name)
	return context2.GetVaultClient(ctx).CreateTransitKey(ctx, name)
}

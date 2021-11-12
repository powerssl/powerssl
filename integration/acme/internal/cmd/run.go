package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"powerssl.dev/common/runner"
	"powerssl.dev/sdk/integration"

	acme "powerssl.dev/integration/acme/internal"
)

func newCmdRun() *cobra.Command {
	cfg := new(acme.Config)
	cmd := runner.RunCmd(&cobra.Command{
		Use:   "run",
		Short: "RunCmd ACME integration",
		Args:  cobra.NoArgs,
	}, cfg, func(ctx context.Context) ([]func() error, func(), error) {
		handler, err := acme.New(cfg.Vault)
		if err != nil {
			return nil, nil, err
		}
		return integration.InitializeACME(ctx, &cfg.Integration, handler)
	})
	return cmd
}

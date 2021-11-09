package cmd

import (
	"context"

	"github.com/spf13/cobra"

	cmdutil "powerssl.dev/common/cmd"
	acme "powerssl.dev/integration/acme/internal"
	"powerssl.dev/sdk/integration"
)

func newCmdRun() *cobra.Command {
	cfg := new(acme.Config)
	cmd := cmdutil.InitAndRun(&cobra.Command{
		Use:   "run",
		Short: "Run ACME integration",
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

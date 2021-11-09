package cmd

import (
	"context"

	"github.com/spf13/cobra"

	cmdutil "powerssl.dev/common/cmd"
	cloudflare "powerssl.dev/integration/cloudflare/internal"
	"powerssl.dev/sdk/integration"
)

func newCmdRun() *cobra.Command {
	cfg := new(cloudflare.Config)
	cmd := cmdutil.InitAndRun(&cobra.Command{
		Use:   "run",
		Short: "Run CloudFlare integration",
		Args:  cobra.NoArgs,
	}, cfg, func(ctx context.Context) ([]func() error, func(), error) {
		handler := cloudflare.New()
		return integration.InitializeDNS(ctx, &cfg.Integration, handler)
	})
	return cmd
}

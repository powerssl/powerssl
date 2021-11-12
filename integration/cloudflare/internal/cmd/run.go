package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"powerssl.dev/common/runner"
	cloudflare "powerssl.dev/integration/cloudflare/internal"
	"powerssl.dev/sdk/integration"
)

func newCmdRun() *cobra.Command {
	cfg := new(cloudflare.Config)
	cmd := runner.RunCmd(&cobra.Command{
		Use:   "run",
		Short: "RunCmd CloudFlare integration",
		Args:  cobra.NoArgs,
	}, cfg, func(ctx context.Context) ([]func() error, func(), error) {
		handler := cloudflare.New()
		return integration.InitializeDNS(ctx, &cfg.Integration, handler)
	})
	return cmd
}

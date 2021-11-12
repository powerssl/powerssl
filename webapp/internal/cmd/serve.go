package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"powerssl.dev/common/runner"

	"powerssl.dev/webapp/internal"
)

func newCmdServe() *cobra.Command {
	cfg := new(internal.Config)
	cmd := runner.RunCmd(&cobra.Command{
		Use:   "serve",
		Short: "Serve the WebApp",
		Args:  cobra.NoArgs,
	}, cfg, func(ctx context.Context) ([]func() error, func(), error) {
		return internal.Initialize(ctx, cfg)
	})
	return cmd
}

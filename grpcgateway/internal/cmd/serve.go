package cmd

import (
	"context"

	"github.com/spf13/cobra"

	cmdutil "powerssl.dev/common/cmd"

	"powerssl.dev/grpcgateway/internal"
)

func newCmdServe() *cobra.Command {
	cfg := new(internal.Config)
	cmd := cmdutil.InitAndRun(&cobra.Command{
		Use:   "serve",
		Short: "Serve the gRPC Gateway",
		Args:  cobra.NoArgs,
	}, cfg, func(ctx context.Context) ([]func() error, func(), error) {
		return internal.Initialize(ctx, cfg)
	})
	return cmd
}

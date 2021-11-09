package cmd

import (
	"context"

	"github.com/spf13/cobra"

	cmdutil "powerssl.dev/common/cmd"

	"powerssl.dev/worker/internal"
)

func newCmdRun() *cobra.Command {
	cfg := new(internal.Config)
	cmd := cmdutil.InitAndRun(&cobra.Command{
		Use:   "run",
		Short: "Run the Worker",
		Args:  cobra.NoArgs,
	}, cfg, func(ctx context.Context) ([]func() error, func(), error) {
		return internal.Initialize(ctx, cfg)
	})
	return cmd
}

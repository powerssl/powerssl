package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"powerssl.dev/common/runner"

	"powerssl.dev/agent/internal"
)

func newCmdRun() *cobra.Command {
	cfg := new(internal.Config)
	cmd := runner.RunCmd(&cobra.Command{
		Use:   "run",
		Short: "RunCmd the Agent",
		Args:  cobra.NoArgs,
	}, cfg, func(ctx context.Context) ([]func() error, func(), error) {
		return internal.Initialize(ctx, cfg)
	})
	return cmd
}

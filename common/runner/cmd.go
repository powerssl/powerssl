package runner // import "powerssl.dev/common/runner"

import (
	"context"

	"github.com/spangenberg/snakecharmer"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func RunCmd(cmd *cobra.Command, cfg snakecharmer.Config, f func(ctx context.Context) ([]func() error, func(), error)) *cobra.Command {
	cmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		if err := viper.Unmarshal(&cfg); err != nil {
			return err
		}
		return snakecharmer.Validate(cfg)
	}
	cmd.Run = snakecharmer.HandleError(func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		runner, ctx := New(ctx)
		fn, cleanup, err := f(ctx)
		if err != nil {
			return err
		}
		defer cleanup()
		x, fn := fn[0], fn[1:]
		return runner.Run(x, fn...)
	})
	snakecharmer.GenerateFlags(cmd.Flags(), cfg)
	return cmd
}

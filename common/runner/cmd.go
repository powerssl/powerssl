package runner // import "powerssl.dev/common/runner"

import (
	"context"

	"github.com/spangenberg/snakecharmer"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func RunCmd(cmd *cobra.Command, cfg snakecharmer.Config, f func(ctx context.Context) ([]func() error, func(), error)) *cobra.Command {
	cmd.PreRunE = snakecharmer.Validate(func(cmd *cobra.Command, args []string) (snakecharmer.Config, error) {
		if err := viper.Unmarshal(&cfg); err != nil {
			return nil, err
		}
		return cfg, nil
	})
	cmd.Run = snakecharmer.HandleError(func(cmd *cobra.Command, args []string) error {
		return Run(f)
	})
	return snakecharmer.GenerateFlags(cmd, cfg)
}

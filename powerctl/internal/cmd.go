package internal

import (
	"context"

	"github.com/spangenberg/snakecharmer"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"powerssl.dev/sdk/apiserver"
)

func CmdWithClient(cmd *cobra.Command, f func(ctx context.Context, client *apiserver.Client, cmd *cobra.Command, args []string) error) *cobra.Command {
	var cleanup func()
	var client *apiserver.Client
	var ctx context.Context
	cfg := new(Config)
	cmd.PreRunE = func(cmd *cobra.Command, args []string) (err error) {
		if err = viper.Unmarshal(&cfg); err != nil {
			return err
		}
		if err = snakecharmer.Validate(cfg); err != nil {
			return err
		}
		ctx = context.Background()
		client, cleanup, err = newClient(ctx, cfg)
		return err
	}
	cmd.Run = snakecharmer.HandleError(func(cmd *cobra.Command, args []string) (err error) {
		return f(ctx, client, cmd, args)
	})
	cmd.PostRun = func(cmd *cobra.Command, args []string) {
		cleanup()
	}
	return cmd
}

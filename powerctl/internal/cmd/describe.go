package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"powerssl.dev/sdk/apiserver"

	"powerssl.dev/powerctl/internal"
	"powerssl.dev/powerctl/internal/resource"
)

func newCmdDescribe() *cobra.Command {
	return internal.CmdWithClient(&cobra.Command{
		Use:   "describe",
		Short: "Describe resource",
		Args:  cobra.RangeArgs(1, 2),
	}, func(ctx context.Context, client *apiserver.Client, cmd *cobra.Command, args []string) error {
		res, err := resource.ResourceFromArgs(args)
		if err != nil {
			return err
		}
		if res, err = res.Get(ctx, client); err != nil {
			return err
		}
		return res.Describe(ctx, client, cmd.OutOrStdout())
	})
}

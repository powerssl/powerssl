package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"powerssl.dev/sdk/apiserver"

	"powerssl.dev/powerctl/internal"
	"powerssl.dev/powerctl/internal/resource"
)

func newCmdDelete() *cobra.Command {
	return internal.CmdWithClient(&cobra.Command{
		Use:   "delete",
		Short: "Delete resource",
		Args:  cobra.MinimumNArgs(1),
	}, func(ctx context.Context, client *apiserver.Client, cmd *cobra.Command, args []string) error {
		resources, err := resource.ResourcesFromArgs(args)
		if err != nil {
			return err
		}
		for _, res := range resources {
			if err = res.Delete(ctx, client); err != nil {
				return err
			}
		}
		return nil
	})
}

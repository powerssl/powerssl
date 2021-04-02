package cmd

import (
	"github.com/spf13/cobra"

	cmdutil "powerssl.dev/common/cmd"
	"powerssl.dev/sdk/apiserver"

	"powerssl.dev/powerctl/internal"
	"powerssl.dev/powerctl/internal/resource"
)

func newCmdDelete() *cobra.Command {
	var client *apiserver.Client

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete resource",
		Args:  cobra.MinimumNArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) (err error) {
			client, err = internal.NewGRPCClient()
			return err

		},
		Run: cmdutil.HandleError(func(cmd *cobra.Command, args []string) (err error) {
			var resources []*resource.Resource
			if resources, err = resource.ResourcesFromArgs(args); err != nil {
				return err
			}
			for _, res := range resources {
				if err = res.Delete(client); err != nil {
					return err
				}
			}
			return nil
		}),
	}

	return cmd
}

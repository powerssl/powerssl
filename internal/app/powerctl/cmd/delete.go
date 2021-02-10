package cmd

import (
	"github.com/spf13/cobra"
	cmdutil "powerssl.dev/powerssl/internal/pkg/cmd"

	"powerssl.dev/powerssl/internal/app/powerctl"
	"powerssl.dev/powerssl/internal/app/powerctl/resource"
	apiserverclient "powerssl.dev/powerssl/pkg/apiserver/client"
)

func newCmdDelete() *cobra.Command {
	var client *apiserverclient.GRPCClient

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete resource",
		Args:  cobra.MinimumNArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) (err error) {
			client, err = powerctl.NewGRPCClient()
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

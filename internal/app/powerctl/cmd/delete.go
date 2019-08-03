package cmd

import (
	"github.com/spf13/cobra"

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
		RunE: func(cmd *cobra.Command, args []string) error {
			resources, err := resource.ResourcesFromArgs(args)
			if err != nil {
				return err
			}
			for _, resource := range resources {
				if err := resource.Delete(client); err != nil {
					return err
				}
			}
			return nil
		},
	}

	return cmd
}

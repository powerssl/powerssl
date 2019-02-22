package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"

	"powerssl.io/internal/app/powerctl"
	"powerssl.io/internal/app/powerctl/resource"
	apiserverclient "powerssl.io/pkg/apiserver/client"
)

func newCmdGet() *cobra.Command {
	var client *apiserverclient.GRPCClient

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get resource",
		Args:  cobra.MinimumNArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) (err error) {
			client, err = powerctl.NewGRPCClient()
			return err
		},
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			var resources []*resource.Resource
			if len(args) == 1 && !strings.Contains(args[0], "/") {
				kinds := strings.Split(args[0], ",")
				if len(kinds) == 1 && kinds[0] == "all" {
					kinds = resource.Kinds()
				}
				for _, kind := range kinds {
					resourceHandler, err := resource.ResourceHandlerByKind(kind)
					if err != nil {
						return err
					}
					res, err := resourceHandler.List(client)
					if err != nil {
						return err
					}
					resources = append(resources, res...)
				}
			} else {
				resources, err = resource.ResourcesFromArgs(args)
				if err != nil {
					return err
				}
				for i, res := range resources {
					if resources[i], err = res.Get(client); err != nil {
						return err
					}
				}
			}
			if len(resources) > 1 {
				return resource.FormatResource(resources, os.Stdout)
			} else if len(resources) == 1 {
				return resource.FormatResource(resources[0], os.Stdout)
			}
			return nil
		},
	}

	return cmd
}

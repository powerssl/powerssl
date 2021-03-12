package cmd

import (
	"strings"

	"github.com/spf13/cobra"

	cmdutil "powerssl.dev/common/cmd"
	"powerssl.dev/powerctl/internal"
	"powerssl.dev/powerctl/internal/resource"
	"powerssl.dev/sdk/apiserver"
)

func newCmdGet() *cobra.Command {
	var client *apiserver.Client

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get resource",
		Args:  cobra.MinimumNArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) (err error) {
			client, err = internal.NewGRPCClient()
			return err
		},
		Run: cmdutil.HandleError(func(cmd *cobra.Command, args []string) (err error) {
			var resources []*resource.Resource
			if len(args) == 1 && !strings.Contains(args[0], "/") {
				kinds := strings.Split(args[0], ",")
				if len(kinds) == 1 && kinds[0] == "all" {
					kinds = resource.Kinds()
				}
				for _, kind := range kinds {
					var resourceHandler resource.Handler
					if resourceHandler, err = resource.ResourceHandlerByKind(kind); err != nil {
						return err
					}
					var res []*resource.Resource
					if res, err = resourceHandler.List(client); err != nil {
						return err
					}
					resources = append(resources, res...)
				}
			} else {
				if resources, err = resource.ResourcesFromArgs(args); err != nil {
					return err
				}
				for i, res := range resources {
					if resources[i], err = res.Get(client); err != nil {
						return err
					}
				}
			}
			if len(resources) > 1 {
				return resource.FormatResource(resources, cmd.OutOrStdout())
			} else if len(resources) == 1 {
				return resource.FormatResource(resources[0], cmd.OutOrStdout())
			}
			return nil
		}),
	}

	return cmd
}

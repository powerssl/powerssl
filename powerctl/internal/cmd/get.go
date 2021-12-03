package cmd

import (
	"context"
	"strings"

	"github.com/spf13/cobra"

	"powerssl.dev/sdk/apiserver"

	"powerssl.dev/powerctl/internal"
	"powerssl.dev/powerctl/internal/resource"
)

func newCmdGet() *cobra.Command {
	return internal.CmdWithClient(&cobra.Command{
		Use:   "get",
		Short: "Get resource",
		Args:  cobra.MinimumNArgs(1),
		Example: `  powerctl get acmeserver
  powerctl get as
  powerctl get acmeserver/7a88f38c-676c-449a-aeac-c78bdd9a78be
  powerctl get c 1e5403f2-5f3d-4c75-b4eb-05ba3c1a3ceb
  powerctl get as/7a88f38c-676c-449a-aeac-c78bdd9a78be c/1e5403f2-5f3d-4c75-b4eb-05ba3c1a3ceb
`,
	}, func(ctx context.Context, client *apiserver.Client, cmd *cobra.Command, args []string) error {
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
				var res []*resource.Resource
				if res, err = resourceHandler.List(ctx, client); err != nil {
					return err
				}
				resources = append(resources, res...)
			}
		} else {
			var err error
			if resources, err = resource.ResourcesFromArgs(args); err != nil {
				return err
			}
			for i, res := range resources {
				if resources[i], err = res.Get(ctx, client); err != nil {
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
	})
}

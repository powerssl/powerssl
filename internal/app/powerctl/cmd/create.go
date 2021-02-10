package cmd

import (
	cmdutil "powerssl.dev/powerssl/internal/pkg/cmd"

	"github.com/spf13/cobra"

	"powerssl.dev/powerssl/internal/app/powerctl"
	"powerssl.dev/powerssl/internal/app/powerctl/resource"
	apiserverclient "powerssl.dev/powerssl/pkg/apiserver/client"
)

func newCmdCreate() *cobra.Command {
	var client *apiserverclient.GRPCClient
	var filename string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create resource",
		Args:  cobra.NoArgs,
		Example: `  # Create a certificate using the data in certificate.json.
  powerctl create -f ./certificate.json

  # Create a certificate based on the JSON passed into stdin.
  cat certificate.json | powerctl create -f -`,
		PreRunE: func(cmd *cobra.Command, args []string) (err error) {
			client, err = powerctl.NewGRPCClient()
			return err
		},
		Run: cmdutil.HandleError(func(cmd *cobra.Command, args []string) (err error) {
			var resources []*resource.Resource
			if resources, err = resource.ResourcesFromFile(filename); err != nil {
				return err
			}
			for i, res := range resources {
				if resources[i], err = res.Create(client); err != nil {
					return err
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

	cmd.Flags().StringVarP(&filename, "filename", "f", "", "Filename to file to use to create the resources")

	return cmd
}

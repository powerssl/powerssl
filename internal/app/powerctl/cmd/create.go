package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"powerssl.io/internal/app/powerctl"
	"powerssl.io/internal/app/powerctl/resource"
	apiserverclient "powerssl.io/pkg/apiserver/client"
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
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			resources, err := resource.ResourcesFromFile(filename)
			if err != nil {
				return err
			}
			for i, res := range resources {
				if resources[i], err = res.Create(client); err != nil {
					return err
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

	cmd.Flags().StringVarP(&filename, "filename", "f", "", "Filename to file to use to create the resources")

	return cmd
}
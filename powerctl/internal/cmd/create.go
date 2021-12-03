package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"powerssl.dev/sdk/apiserver"

	"powerssl.dev/powerctl/internal"
	"powerssl.dev/powerctl/internal/resource"
)

func newCmdCreate() *cobra.Command {
	var filename string

	cmd := internal.CmdWithClient(&cobra.Command{
		Use:   "create",
		Short: "Create resource",
		Args:  cobra.NoArgs,
		Example: `  # Create a certificate using the data in certificate.json.
  powerctl create -f ./certificate.json

  # Create a certificate based on the JSON passed into stdin.
  cat certificate.json | powerctl create -f -`,
	}, func(ctx context.Context, client *apiserver.Client, cmd *cobra.Command, args []string) error {
		resources, err := resource.ResourcesFromFile(filename)
		if err != nil {
			return err
		}
		for i, res := range resources {
			if resources[i], err = res.Create(ctx, client); err != nil {
				return err
			}
		}
		if len(resources) > 1 {
			return resource.FormatResource(resources, cmd.OutOrStdout())
		} else if len(resources) == 1 {
			return resource.FormatResource(resources[0], cmd.OutOrStdout())
		}
		return nil
	})

	cmd.Flags().StringVarP(&filename, "filename", "f", "", "Filename to file to use to create the resources")

	cmd.AddCommand(resource.NewCmdCreateACMEAccount())
	cmd.AddCommand(resource.NewCmdCreateACMEServer())
	cmd.AddCommand(resource.NewCmdCreateCertificate())
	cmd.AddCommand(resource.NewCmdCreateUser())

	return cmd
}

package cmd

import (
	"github.com/spf13/cobra"

	cmdutil "powerssl.dev/common/cmd"
	"powerssl.dev/powerctl/internal"
	"powerssl.dev/powerctl/internal/resource"
	"powerssl.dev/sdk/apiserver"
)

func newCmdCreate() *cobra.Command {
	var client *apiserver.Client
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
			client, err = internal.NewGRPCClient()
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

	cmd.AddCommand(resource.NewCmdCreateACMEAccount())
	cmd.AddCommand(resource.NewCmdCreateACMEServer())
	cmd.AddCommand(resource.NewCmdCreateCertificate())
	cmd.AddCommand(resource.NewCmdCreateUser())

	return cmd
}

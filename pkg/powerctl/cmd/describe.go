package cmd

import (
	"os"

	"github.com/spf13/cobra"

	apiserverclient "powerssl.io/pkg/apiserver/client"
	"powerssl.io/pkg/powerctl"
	"powerssl.io/pkg/powerctl/resource"
)

func newCmdDescribe() *cobra.Command {
	var client *apiserverclient.GRPCClient

	cmd := &cobra.Command{
		Use:   "describe",
		Short: "Describe resource",
		Args:  cobra.RangeArgs(1, 2),
		PreRunE: func(cmd *cobra.Command, args []string) (err error) {
			client, err = powerctl.NewGRPCClient()
			return err

		},
		RunE: func(cmd *cobra.Command, args []string) error {
			res, err := resource.ResourceFromArgs(args)
			if err != nil {
				return err
			}
			if res, err = res.Get(client); err != nil {
				return err
			}
			return res.Describe(client, os.Stdout)
		},
	}

	return cmd
}

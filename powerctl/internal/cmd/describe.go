package cmd

import (
	"github.com/spf13/cobra"

	cmdutil "powerssl.dev/common/cmd"
	"powerssl.dev/powerctl/internal"
	"powerssl.dev/powerctl/internal/resource"
	apiserverclient "powerssl.dev/sdk/apiserver/client"
)

func newCmdDescribe() *cobra.Command {
	var client *apiserverclient.GRPCClient

	cmd := &cobra.Command{
		Use:   "describe",
		Short: "Describe resource",
		Args:  cobra.RangeArgs(1, 2),
		PreRunE: func(cmd *cobra.Command, args []string) (err error) {
			client, err = internal.NewGRPCClient()
			return err
		},
		Run: cmdutil.HandleError(func(cmd *cobra.Command, args []string) (err error) {
			var res *resource.Resource
			if res, err = resource.ResourceFromArgs(args); err != nil {
				return err
			}
			if res, err = res.Get(client); err != nil {
				return err
			}
			return res.Describe(client, cmd.OutOrStdout())
		}),
	}

	return cmd
}

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Describe resource",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := NewGRPCClient()
		if err != nil {
			er(err)
		}
		resource, err := resourceFromArgs(args)
		if err != nil {
			er(err)
		}
		if resource, err = resource.Get(client); err != nil {
			er(err)
		}
		if err := resource.Describe(client, os.Stdout); err != nil {
			er(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(describeCmd)
}

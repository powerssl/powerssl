package cmd

import "github.com/spf13/cobra"

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete resource",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := NewGRPCClient()
		if err != nil {
			er(err)
		}
		resources, err := resourcesFromArgs(args)
		if err != nil {
			er(err)
		}
		for _, resource := range resources {
			if err := resource.Delete(client); err != nil {
				er(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

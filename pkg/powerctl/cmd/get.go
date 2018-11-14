package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get resource",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := NewGRPCClient()
		if err != nil {
			er(err)
		}
		var resources []*Resource
		if len(args) == 1 && !strings.Contains(args[0], "/") {
			resourceHandler, err := Resources.Get(args[0])
			if err != nil {
				er(err)
			}
			resources, err = resourceHandler.List(client)
		} else {
			resources, err = resourcesFromArgs(args)
			if err != nil {
				er(err)
			}
			for i, resource := range resources {
				if resources[i], err = resource.Get(client); err != nil {
					er(err)
				}
			}
		}
		if len(resources) > 1 {
			pr(resources)
		} else if len(resources) == 1 {
			pr(resources[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

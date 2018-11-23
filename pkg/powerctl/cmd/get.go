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
			kinds := strings.Split(args[0], ",")
			for _, kind := range kinds {
				resourceHandler, err := Resources.Get(kind)
				if err != nil {
					er(err)
				}
				res, err := resourceHandler.List(client)
				if err != nil {
					er(err)
				}
				resources = append(resources, res...)
			}
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

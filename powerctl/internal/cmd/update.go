package cmd

import "github.com/spf13/cobra"

func newCmdUpdate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update resource",
	}

	return cmd
}
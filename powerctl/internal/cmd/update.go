package cmd

import (
	"github.com/spf13/cobra"
)

func newCmdUpdate() *cobra.Command {
	return &cobra.Command{
		Use:   "update",
		Short: "Update resource",
	}
}

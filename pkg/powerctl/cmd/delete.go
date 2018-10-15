package cmd

import "github.com/spf13/cobra"

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete resource",
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
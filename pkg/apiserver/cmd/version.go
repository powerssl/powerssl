package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version: version.Info{Major:\"0\", Minor:\"1\"}")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

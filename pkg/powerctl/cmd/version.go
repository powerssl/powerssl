package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the client and server version information",
	Long: `Print the client and server version information

Examples:
  # Print the client and server versions
  powerctl version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Client Version: version.Info{Major:\"0\", Minor:\"1\"}")
		fmt.Println("Server Version: version.Info{Major:\"\", Minor:\"\"}")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

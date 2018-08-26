package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Display one or many resources",
	Long: `Display one or many resources.

Examples:
  # List all certificate authorities
  powerctl get certificateauthorities`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
		fmt.Println(viper.GetString("server"))
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

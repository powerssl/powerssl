package cmd

import (
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a resource.",
	Long: `Create a resource.

Available Commands:
  certificateauthority Create a CertificateAuthority`,
	Args: cobra.ExactArgs(1),
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("create called")
	//	fmt.Println(viper.GetString("server"))
	//},
}

func init() {
	rootCmd.AddCommand(createCmd)
}

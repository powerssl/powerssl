package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"powerssl.io/pkg/agent"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the Agent",
	Run: func(cmd *cobra.Command, args []string) {
		addr := viper.GetString("addr")
		authToken := viper.GetString("auth-token")
		caFile := viper.GetString("ca-file")
		insecure := viper.GetBool("insecure")
		insecureSkipTLSVerify := viper.GetBool("insecure-skip-tls-verify")
		serverNameOverride := viper.GetString("server-name-override")

		ok := true
		if addr == "" {
			ok = false
			fmt.Println("Provide addr")
		}
		if authToken == "" {
			ok = false
			fmt.Println("Provide auth-token")
		}
		if !insecure && !insecureSkipTLSVerify && caFile == "" {
			ok = false
			fmt.Println("Provide ca-file")
		}
		if !ok {
			os.Exit(1)
		}

		agent.Run(addr, caFile, serverNameOverride, insecure, insecureSkipTLSVerify, authToken)
	},
}

func init() {
	runCmd.Flags().BoolP("insecure", "", false, "Use insecure communication")
	runCmd.Flags().BoolP("insecure-skip-tls-verify", "", false, "Accepts any certificate presented by the server and any host name in that certificate")
	runCmd.Flags().StringP("addr", "", "", "GRPC address of API")
	runCmd.Flags().StringP("auth-token", "", "", "Auth token")
	runCmd.Flags().StringP("ca-file", "", "", "Certificate authority file")
	runCmd.Flags().StringP("server-name-override", "", "", "It will override the virtual host name of authority")

	viper.BindPFlag("addr", runCmd.Flags().Lookup("addr"))
	viper.BindPFlag("auth-token", runCmd.Flags().Lookup("auth-token"))
	viper.BindPFlag("ca-file", runCmd.Flags().Lookup("ca-file"))
	viper.BindPFlag("insecure", runCmd.Flags().Lookup("insecure"))
	viper.BindPFlag("insecure-skip-tls-verify", runCmd.Flags().Lookup("insecure-skip-tls-verify"))
	viper.BindPFlag("server-name-override", runCmd.Flags().Lookup("server-name-override"))

	rootCmd.AddCommand(runCmd)
}

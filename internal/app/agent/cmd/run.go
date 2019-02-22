package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"powerssl.io/internal/app/agent"
)

func newCmdRun() *cobra.Command {
	cmd := &cobra.Command{
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

	cmd.Flags().BoolP("insecure", "", false, "Use insecure communication")
	cmd.Flags().BoolP("insecure-skip-tls-verify", "", false, "Accepts any certificate presented by the server and any host name in that certificate")
	cmd.Flags().StringP("addr", "", "", "GRPC address of API")
	cmd.Flags().StringP("auth-token", "", "", "Auth token")
	cmd.Flags().StringP("ca-file", "", "", "Certificate authority file")
	cmd.Flags().StringP("server-name-override", "", "", "It will override the virtual host name of authority")

	viper.BindPFlag("addr", cmd.Flags().Lookup("addr"))
	viper.BindPFlag("auth-token", cmd.Flags().Lookup("auth-token"))
	viper.BindPFlag("ca-file", cmd.Flags().Lookup("ca-file"))
	viper.BindPFlag("insecure", cmd.Flags().Lookup("insecure"))
	viper.BindPFlag("insecure-skip-tls-verify", cmd.Flags().Lookup("insecure-skip-tls-verify"))
	viper.BindPFlag("server-name-override", cmd.Flags().Lookup("server-name-override"))

	return cmd
}
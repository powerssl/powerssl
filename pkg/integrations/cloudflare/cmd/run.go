package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"powerssl.io/pkg/integration"
	"powerssl.io/pkg/integrations/cloudflare"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run ACME integration",
	Run: func(cmd *cobra.Command, args []string) {
		certFile := viper.GetString("ca-file")
		addr := viper.GetString("addr")
		insecure := viper.GetBool("insecure")
		insecureSkipTLSVerify := viper.GetBool("insecure-skip-tls-verify")
		serverNameOverride := viper.GetString("server-name-override")

		ok := true
		if addr == "" {
			ok = false
			fmt.Println("Provide addr")
		}
		if !insecure && !insecureSkipTLSVerify && certFile == "" {
			ok = false
			fmt.Println("Provide ca-file")
		}
		if !ok {
			os.Exit(1)
		}

		integ := integration.New(addr, certFile, serverNameOverride, insecure, insecureSkipTLSVerify, integration.KindDNS, "cloudflare")
		integ.Run(cloudflare.New())
	},
}

func init() {
	runCmd.Flags().BoolP("insecure", "", false, "Use insecure communication")
	runCmd.Flags().BoolP("insecure-skip-tls-verify", "", false, "Accepts any certificate presented by the server and any host name in that certificate")
	runCmd.Flags().StringP("addr", "", "", "GRPC address of Controller")
	runCmd.Flags().StringP("ca-file", "", "", "Certificate authority file")
	runCmd.Flags().StringP("server-name-override", "", "", "It will override the virtual host name of authority")

	viper.BindPFlag("addr", runCmd.Flags().Lookup("addr"))
	viper.BindPFlag("ca-file", runCmd.Flags().Lookup("ca-file"))
	viper.BindPFlag("insecure", runCmd.Flags().Lookup("insecure"))
	viper.BindPFlag("insecure-skip-tls-verify", runCmd.Flags().Lookup("insecure-skip-tls-verify"))
	viper.BindPFlag("server-name-override", runCmd.Flags().Lookup("server-name-override"))

	rootCmd.AddCommand(runCmd)
}

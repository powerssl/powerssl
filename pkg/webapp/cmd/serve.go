package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"powerssl.io/pkg/webapp"
)

func newCmdServe() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Serve the API",
		Run: func(cmd *cobra.Command, args []string) {
			addr := viper.GetString("addr")
			var metricsAddr string
			if !viper.GetBool("no-metrics") {
				metricsAddr = viper.GetString("metrics-addr")
			}
			insecure := viper.GetBool("insecure")
			tlsCertFile := viper.GetString("tls.cert-file")
			tlsPrivateKeyFile := viper.GetString("tls.private-key-file")

			ok := true
			if addr == "" {
				ok = false
				fmt.Println("Provide addr")
			}
			if !insecure && tlsCertFile == "" {
				ok = false
				fmt.Println("Provide tls-cert-file")
			}
			if !insecure && tlsPrivateKeyFile == "" {
				ok = false
				fmt.Println("Provide tls-private-key-file")
			}
			if !ok {
				os.Exit(1)
			}

			webapp.Run(addr, tlsCertFile, tlsPrivateKeyFile, insecure, metricsAddr)
		},
	}

	cmd.Flags().BoolP("insecure", "", false, "Do not use TLS for the server")
	cmd.Flags().BoolP("no-metrics", "", false, "Do not serve metrics")
	cmd.Flags().StringP("addr", "", ":8080", "Addr")
	cmd.Flags().StringP("metrics-addr", "", ":9090", "HTTP Addr")
	cmd.Flags().StringP("tls-cert-file", "", "", "File containing the default x509 Certificate.")
	cmd.Flags().StringP("tls-private-key-file", "", "", "File containing the default x509 private key matching --tls-cert-file.")

	viper.BindPFlag("addr", cmd.Flags().Lookup("addr"))
	viper.BindPFlag("auth-token", cmd.Flags().Lookup("auth-token"))
	viper.BindPFlag("insecure", cmd.Flags().Lookup("insecure"))
	viper.BindPFlag("metrics-addr", cmd.Flags().Lookup("metrics-addr"))
	viper.BindPFlag("no-metrics", cmd.Flags().Lookup("no-metrics"))
	viper.BindPFlag("tls.cert-file", cmd.Flags().Lookup("tls-cert-file"))
	viper.BindPFlag("tls.private-key-file", cmd.Flags().Lookup("tls-private-key-file"))

	return cmd
}

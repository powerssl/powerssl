package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"powerssl.io/pkg/signer"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the API",
	Run: func(cmd *cobra.Command, args []string) {
		addr := viper.GetString("addr")
		var httpAddr string
		if !viper.GetBool("no-metrics") {
			httpAddr = viper.GetString("metrics-addr")
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

		signer.Run(addr, tlsCertFile, tlsPrivateKeyFile, insecure, httpAddr)
	},
}

func init() {
	serveCmd.Flags().BoolP("insecure", "", false, "Do not use TLS for the server")
	serveCmd.Flags().BoolP("no-metrics", "", false, "Do not serve metrics")
	serveCmd.Flags().StringP("addr", "", ":8080", "GRPC Addr")
	serveCmd.Flags().StringP("metrics-addr", "", ":9090", "HTTP Addr")
	serveCmd.Flags().StringP("tls-cert-file", "", "", "File containing the default x509 Certificate for GRPC.")
	serveCmd.Flags().StringP("tls-private-key-file", "", "", "File containing the default x509 private key matching --tls-cert-file.")

	viper.BindPFlag("addr", serveCmd.Flags().Lookup("addr"))
	viper.BindPFlag("insecure", serveCmd.Flags().Lookup("insecure"))
	viper.BindPFlag("metrics-addr", serveCmd.Flags().Lookup("metrics-addr"))
	viper.BindPFlag("no-metrics", serveCmd.Flags().Lookup("no-metrics"))
	viper.BindPFlag("tls.cert-file", serveCmd.Flags().Lookup("tls-cert-file"))
	viper.BindPFlag("tls.private-key-file", serveCmd.Flags().Lookup("tls-private-key-file"))

	rootCmd.AddCommand(serveCmd)
}

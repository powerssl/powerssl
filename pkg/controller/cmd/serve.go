package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"powerssl.io/pkg/controller"
)

func newCmdServe() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Serve the API",
		Run: func(cmd *cobra.Command, args []string) {
			addr := viper.GetString("addr")
			apiserverAddr := viper.GetString("apiserver.addr")
			apiserverAuthToken := viper.GetString("apiserver.auth-token")
			apiserverInsecure := viper.GetBool("apiserver.insecure")
			apiserverInsecureSkipTLSVerify := viper.GetBool("apiserver.insecure-skip-tls-verify")
			apiserverServerNameOverride := viper.GetString("apiserver.server-name-override")
			caFile := viper.GetString("ca-file")
			commonName := viper.GetString("common-name")
			insecure := viper.GetBool("insecure")
			jwksURL := viper.GetString("jwks-url")
			tlsCertFile := viper.GetString("tls.cert-file")
			tlsPrivateKeyFile := viper.GetString("tls.private-key-file")
			vaultToken := viper.GetString("vault.token")
			vaultURL := viper.GetString("vault.url")

			var metricsAddr string
			if !viper.GetBool("no-metrics") {
				metricsAddr = viper.GetString("metrics-addr")
			}
			var tracer string
			if !viper.GetBool("no-tracing") {
				tracer = viper.GetString("tracer")
			}

			ok := true
			if addr == "" {
				ok = false
				fmt.Println("Provide addr")
			}
			if !insecure && commonName == "" && (tlsCertFile == "" || tlsPrivateKeyFile == "") {
				ok = false
				fmt.Println("Provide common-name or tls-cert-file and tls-private-key")
			}
			if jwksURL == "" {
				ok = false
				fmt.Println("Provide jwks-url")
			}
			if apiserverAddr == "" {
				ok = false
				fmt.Println("Provide apiserver-addr")
			}
			if !apiserverInsecure && !apiserverInsecureSkipTLSVerify && caFile == "" {
				ok = false
				fmt.Println("Provide ca-file")
			}
			if !ok {
				os.Exit(1)
			}

			controller.Run(addr, commonName, vaultURL, vaultToken, tlsCertFile, tlsPrivateKeyFile, insecure, metricsAddr, tracer, apiserverAddr, caFile, apiserverServerNameOverride, apiserverInsecure, apiserverInsecureSkipTLSVerify, jwksURL, apiserverAuthToken)
		},
	}

	cmd.Flags().BoolP("apiserver-insecure", "", false, "Use insecure communication")
	cmd.Flags().BoolP("apiserver-insecure-skip-tls-verify", "", false, "Accepts any certificate presented by the server and any host name in that certificate")
	cmd.Flags().BoolP("insecure", "", false, "Do not use TLS for the server")
	cmd.Flags().BoolP("no-metrics", "", false, "Do not serve metrics")
	cmd.Flags().BoolP("no-tracing", "", false, "Do not trace")
	cmd.Flags().StringP("addr", "", ":8080", "GRPC Addr")
	cmd.Flags().StringP("apiserver-addr", "", "", "GRPC address of API server")
	cmd.Flags().StringP("apiserver-auth-token", "", "", "API server authentication token")
	cmd.Flags().StringP("apiserver-server-name-override", "", "", "It will override the virtual host name of authority")
	cmd.Flags().StringP("ca-file", "", "", "Certificate authority file")
	cmd.Flags().StringP("common-name", "", "", "API Server common name")
	cmd.Flags().StringP("jwks-url", "", "", "JWKS URL")
	cmd.Flags().StringP("metrics-addr", "", ":9090", "HTTP Addr")
	cmd.Flags().StringP("tls-cert-file", "", "", "File containing the default x509 Certificate for GRPC.")
	cmd.Flags().StringP("tls-private-key-file", "", "", "File containing the default x509 private key matching --tls-cert-file.")
	cmd.Flags().StringP("tracer", "", "jaeger", "Tracing implementation")
	cmd.Flags().StringP("vault-token", "", "", "Vault Token")
	cmd.Flags().StringP("vault-url", "", "", "Vault URL")

	viper.BindPFlag("addr", cmd.Flags().Lookup("addr"))
	viper.BindPFlag("apiserver.addr", cmd.Flags().Lookup("apiserver-addr"))
	viper.BindPFlag("apiserver.auth-token", cmd.Flags().Lookup("apiserver-auth-token"))
	viper.BindPFlag("apiserver.insecure", cmd.Flags().Lookup("apiserver-insecure"))
	viper.BindPFlag("apiserver.insecure-skip-tls-verify", cmd.Flags().Lookup("apiserver-insecure-skip-tls-verify"))
	viper.BindPFlag("apiserver.server-name-override", cmd.Flags().Lookup("apiserver-server-name-override"))
	viper.BindPFlag("ca-file", cmd.Flags().Lookup("ca-file"))
	viper.BindPFlag("common-name", cmd.Flags().Lookup("common-name"))
	viper.BindPFlag("insecure", cmd.Flags().Lookup("insecure"))
	viper.BindPFlag("jwks-url", cmd.Flags().Lookup("jwks-url"))
	viper.BindPFlag("metrics-addr", cmd.Flags().Lookup("metrics-addr"))
	viper.BindPFlag("no-metrics", cmd.Flags().Lookup("no-metrics"))
	viper.BindPFlag("no-tracing", cmd.Flags().Lookup("no-tracing"))
	viper.BindPFlag("tls.cert-file", cmd.Flags().Lookup("tls-cert-file"))
	viper.BindPFlag("tls.private-key-file", cmd.Flags().Lookup("tls-private-key-file"))
	viper.BindPFlag("tracer", cmd.Flags().Lookup("tracer"))
	viper.BindPFlag("vault.token", cmd.Flags().Lookup("vault-token"))
	viper.BindPFlag("vault.url", cmd.Flags().Lookup("vault-url"))

	return cmd
}

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"powerssl.io/pkg/controller"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the API",
	Run: func(cmd *cobra.Command, args []string) {
		addr := viper.GetString("addr")
		apiserverAddr := viper.GetString("apiserver.addr")
		apiserverAuthToken := viper.GetString("apiserver.auth-token")
		apiserverCertFile := viper.GetString("apiserver.ca-file")
		apiserverInsecure := viper.GetBool("apiserver.insecure")
		apiserverInsecureSkipTLSVerify := viper.GetBool("apiserver.insecure-skip-tls-verify")
		apiserverServerNameOverride := viper.GetString("apiserver.server-name-override")
		var metricsAddr string
		if !viper.GetBool("no-metrics") {
			metricsAddr = viper.GetString("metrics-addr")
		}
		insecure := viper.GetBool("insecure")
		jwtPublicKeyFile := viper.GetString("jwt.public-key-file")
		tlsCertFile := viper.GetString("tls.cert-file")
		tlsPrivateKeyFile := viper.GetString("tls.private-key-file")
		var tracer string
		if !viper.GetBool("no-tracing") {
			tracer = viper.GetString("tracer")
		}

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
		if jwtPublicKeyFile == "" {
			ok = false
			fmt.Println("Provide jwt-public-key-file")
		}
		if apiserverAddr == "" {
			ok = false
			fmt.Println("Provide apiserver-addr")
		}
		if !apiserverInsecure && !apiserverInsecureSkipTLSVerify && apiserverCertFile == "" {
			ok = false
			fmt.Println("Provide apiserver-ca-file")
		}
		if !ok {
			os.Exit(1)
		}

		controller.Run(addr, tlsCertFile, tlsPrivateKeyFile, insecure, metricsAddr, tracer, apiserverAddr, apiserverCertFile, apiserverServerNameOverride, apiserverInsecure, apiserverInsecureSkipTLSVerify, jwtPublicKeyFile, apiserverAuthToken)
	},
}

func init() {
	serveCmd.Flags().BoolP("apiserver-insecure", "", false, "Use insecure communication")
	serveCmd.Flags().BoolP("apiserver-insecure-skip-tls-verify", "", false, "Accepts any certificate presented by the server and any host name in that certificate")
	serveCmd.Flags().BoolP("insecure", "", false, "Do not use TLS for the server")
	serveCmd.Flags().BoolP("no-metrics", "", false, "Do not serve metrics")
	serveCmd.Flags().BoolP("no-tracing", "", false, "Do not trace")
	serveCmd.Flags().StringP("addr", "", ":8080", "GRPC Addr")
	serveCmd.Flags().StringP("apiserver-addr", "", "", "GRPC address of API server")
	serveCmd.Flags().StringP("apiserver-auth-token", "", "", "API server authentication token")
	serveCmd.Flags().StringP("apiserver-ca-file", "", "", "Certificate authority file")
	serveCmd.Flags().StringP("apiserver-server-name-override", "", "", "It will override the virtual host name of authority")
	serveCmd.Flags().StringP("jwt-public-key-file", "", "", "JWT public key file")
	serveCmd.Flags().StringP("metrics-addr", "", ":9090", "HTTP Addr")
	serveCmd.Flags().StringP("signing-key", "", "", "Signing key")
	serveCmd.Flags().StringP("tls-cert-file", "", "", "File containing the default x509 Certificate for GRPC.")
	serveCmd.Flags().StringP("tls-private-key-file", "", "", "File containing the default x509 private key matching --tls-cert-file.")
	serveCmd.Flags().StringP("tracer", "", "jaeger", "Tracing implementation")

	viper.BindPFlag("addr", serveCmd.Flags().Lookup("addr"))
	viper.BindPFlag("apiserver.addr", serveCmd.Flags().Lookup("apiserver-addr"))
	viper.BindPFlag("apiserver.auth-token", serveCmd.Flags().Lookup("apiserver-auth-token"))
	viper.BindPFlag("apiserver.ca-file", serveCmd.Flags().Lookup("apiserver-ca-file"))
	viper.BindPFlag("apiserver.insecure", serveCmd.Flags().Lookup("apiserver-insecure"))
	viper.BindPFlag("apiserver.insecure-skip-tls-verify", serveCmd.Flags().Lookup("apiserver-insecure-skip-tls-verify"))
	viper.BindPFlag("apiserver.server-name-override", serveCmd.Flags().Lookup("apiserver-server-name-override"))
	viper.BindPFlag("insecure", serveCmd.Flags().Lookup("insecure"))
	viper.BindPFlag("jwt.public-key-file", serveCmd.Flags().Lookup("jwt-public-key-file"))
	viper.BindPFlag("metrics-addr", serveCmd.Flags().Lookup("metrics-addr"))
	viper.BindPFlag("no-metrics", serveCmd.Flags().Lookup("no-metrics"))
	viper.BindPFlag("no-tracing", serveCmd.Flags().Lookup("no-tracing"))
	viper.BindPFlag("signing-key", serveCmd.Flags().Lookup("signing-key"))
	viper.BindPFlag("tls.cert-file", serveCmd.Flags().Lookup("tls-cert-file"))
	viper.BindPFlag("tls.private-key-file", serveCmd.Flags().Lookup("tls-private-key-file"))
	viper.BindPFlag("tracer", serveCmd.Flags().Lookup("tracer"))

	rootCmd.AddCommand(serveCmd)
}

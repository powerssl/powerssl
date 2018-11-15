package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"powerssl.io/pkg/apiserver"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the API",
	Run: func(cmd *cobra.Command, args []string) {
		addr := viper.GetString("addr")
		controllerAddr := viper.GetString("controller.addr")
		controllerCertFile := viper.GetString("controller.ca-file")
		controllerInsecure := viper.GetBool("controller.insecure")
		controllerInsecureSkipTLSVerify := viper.GetBool("controller.insecure-skip-tls-verify")
		controllerServerNameOverride := viper.GetString("controller.server-name-override")
		dbConnection := viper.GetString("db.connection")
		dbDialect := viper.GetString("db.dialect")
		var metricsAddr string
		if !viper.GetBool("no-metrics") {
			metricsAddr = viper.GetString("metrics-addr")
		}
		insecure := viper.GetBool("insecure")
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
		if dbConnection == "" {
			ok = false
			fmt.Println("Provide db-connection")
		}
		if dbDialect == "" {
			ok = false
			fmt.Println("Provide db-dialect")
		}
		if !insecure && tlsCertFile == "" {
			ok = false
			fmt.Println("Provide tls-cert-file")
		}
		if !insecure && tlsPrivateKeyFile == "" {
			ok = false
			fmt.Println("Provide tls-private-key-file")
		}

		if controllerAddr == "" {
			ok = false
			fmt.Println("Provide controller-addr")
		}
		if !controllerInsecure && !controllerInsecureSkipTLSVerify && controllerCertFile == "" {
			ok = false
			fmt.Println("Provide controller-ca-file")
		}
		if !ok {
			os.Exit(1)
		}

		apiserver.Run(addr, tlsCertFile, tlsPrivateKeyFile, insecure, dbDialect, dbConnection, metricsAddr, tracer, controllerAddr, controllerCertFile, controllerServerNameOverride, controllerInsecure, controllerInsecureSkipTLSVerify)
	},
}

func init() {
	serveCmd.Flags().BoolP("controller-insecure", "", false, "Use insecure communication")
	serveCmd.Flags().BoolP("controller-insecure-skip-tls-verify", "", false, "Accepts any certificate presented by the server and any host name in that certificate")
	serveCmd.Flags().BoolP("insecure", "", false, "Do not use TLS for the server")
	serveCmd.Flags().BoolP("no-metrics", "", false, "Do not serve metrics")
	serveCmd.Flags().BoolP("no-tracing", "", false, "Do not trace")
	serveCmd.Flags().StringP("addr", "", ":8080", "GRPC Addr")
	serveCmd.Flags().StringP("auth-token", "", "", "Authentication token")
	serveCmd.Flags().StringP("controller-addr", "", "", "GRPC address of Controller")
	serveCmd.Flags().StringP("controller-ca-file", "", "", "Certificate authority file")
	serveCmd.Flags().StringP("controller-server-name-override", "", "", "It will override the virtual host name of authority")
	serveCmd.Flags().StringP("db-connection", "", "/tmp/powerssl.sqlie3", "DB connection")
	serveCmd.Flags().StringP("db-dialect", "", "sqlite3", "DB Dialect")
	serveCmd.Flags().StringP("metrics-addr", "", ":9090", "HTTP Addr")
	serveCmd.Flags().StringP("tls-cert-file", "", "", "File containing the default x509 Certificate for GRPC.")
	serveCmd.Flags().StringP("tls-private-key-file", "", "", "File containing the default x509 private key matching --tls-cert-file.")
	serveCmd.Flags().StringP("tracer", "", "jaeger", "Tracing implementation")

	viper.BindPFlag("addr", serveCmd.Flags().Lookup("addr"))
	viper.BindPFlag("auth-token", serveCmd.Flags().Lookup("auth-token"))
	viper.BindPFlag("controller.addr", serveCmd.Flags().Lookup("controller-addr"))
	viper.BindPFlag("controller.ca-file", serveCmd.Flags().Lookup("controller-ca-file"))
	viper.BindPFlag("controller.insecure", serveCmd.Flags().Lookup("controller-insecure"))
	viper.BindPFlag("controller.insecure-skip-tls-verify", serveCmd.Flags().Lookup("controller-insecure-skip-tls-verify"))
	viper.BindPFlag("controller.server-name-override", serveCmd.Flags().Lookup("controller-server-name-override"))
	viper.BindPFlag("db.connection", serveCmd.Flags().Lookup("db-connection"))
	viper.BindPFlag("db.dialect", serveCmd.Flags().Lookup("db-dialect"))
	viper.BindPFlag("insecure", serveCmd.Flags().Lookup("insecure"))
	viper.BindPFlag("metrics-addr", serveCmd.Flags().Lookup("metrics-addr"))
	viper.BindPFlag("no-metrics", serveCmd.Flags().Lookup("no-metrics"))
	viper.BindPFlag("no-tracing", serveCmd.Flags().Lookup("no-tracing"))
	viper.BindPFlag("tls.cert-file", serveCmd.Flags().Lookup("tls-cert-file"))
	viper.BindPFlag("tls.private-key-file", serveCmd.Flags().Lookup("tls-private-key-file"))
	viper.BindPFlag("tracer", serveCmd.Flags().Lookup("tracer"))

	rootCmd.AddCommand(serveCmd)
}

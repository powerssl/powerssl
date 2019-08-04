package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"powerssl.dev/powerssl/internal/app/apiserver"
)

func newCmdServe() *cobra.Command {
	var (
		addr                            string
		caFile                          string
		commonName                      string
		controllerAddr                  string
		authToken                       string
		controllerInsecure              bool
		controllerInsecureSkipTLSVerify bool
		controllerServerNameOverride    string
		dbConnection                    string
		dbDialect                       string
		insecure                        bool
		jwksURL                         string
		metricsAddr                     string
		tlsCertFile                     string
		tlsPrivateKeyFile               string
		tracer                          string
		vaultToken                      string
		vaultURL                        string
	)

	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Serve the API",
		PreRun: func(cmd *cobra.Command, args []string) {
			addr = viper.GetString("addr")
			authToken = viper.GetString("auth-token")
			caFile = viper.GetString("ca-file")
			commonName = viper.GetString("common-name")
			controllerAddr = viper.GetString("controller.addr")
			controllerInsecure = viper.GetBool("controller.insecure")
			controllerInsecureSkipTLSVerify = viper.GetBool("controller.insecure-skip-tls-verify")
			controllerServerNameOverride = viper.GetString("controller.server-name-override")
			dbConnection = viper.GetString("db.connection")
			dbDialect = viper.GetString("db.dialect")
			insecure = viper.GetBool("insecure")
			jwksURL = viper.GetString("jwks-url")
			if !viper.GetBool("no-metrics") {
				metricsAddr = viper.GetString("metrics-addr")
			}
			tlsCertFile = viper.GetString("tls.cert-file")
			tlsPrivateKeyFile = viper.GetString("tls.private-key-file")
			if !viper.GetBool("no-tracing") {
				tracer = viper.GetString("tracer")
			}
			vaultToken = viper.GetString("vault.token")
			vaultURL = viper.GetString("vault.url")
		},
		Run: func(cmd *cobra.Command, args []string) {
			apiserver.Run(&apiserver.Config{
				AuthToken: authToken,
				ControllerClientConfig: &apiserver.ControllerClientConfig{
					CAFile:                caFile,
					Addr:                  controllerAddr,
					Insecure:              controllerInsecure,
					InsecureSkipTLSVerify: controllerInsecureSkipTLSVerify,
					ServerNameOverride:    controllerServerNameOverride,
				},
				DBConnection: dbConnection,
				DBDialect:    dbDialect,
				JWKSURL:      jwksURL,
				MetricsAddr:  metricsAddr,
				ServerConfig: &apiserver.ServerConfig{
					Addr:       addr,
					CAFile:     caFile,
					CertFile:   tlsCertFile,
					CommonName: commonName,
					Insecure:   insecure,
					KeyFile:    tlsPrivateKeyFile,
					VaultToken: vaultToken,
					VaultURL:   vaultURL,
				},
				Tracer: tracer,
				VaultClientConfig: &apiserver.VaultClientConfig{
					CAFile: caFile,
					Token:  vaultToken,
					URL:    vaultURL,
				},
			})
		},
	}

	cmd.Flags().BoolP("controller-insecure", "", false, "Use insecure communication")
	cmd.Flags().BoolP("controller-insecure-skip-tls-verify", "", false, "Accepts any certificate presented by the server and any host name in that certificate")
	cmd.Flags().BoolP("insecure", "", false, "Do not use TLS for the server")
	cmd.Flags().BoolP("no-metrics", "", false, "Do not serve metrics")
	cmd.Flags().BoolP("no-tracing", "", false, "Do not trace")
	cmd.Flags().StringP("addr", "", ":8080", "GRPC Addr")
	cmd.Flags().StringP("auth-token", "", "", "Authentication token")
	cmd.Flags().StringP("ca-file", "", "", "Certificate authority file")
	cmd.Flags().StringP("common-name", "", "", "API Server common name")
	cmd.Flags().StringP("controller-addr", "", "", "GRPC address of Controller")
	cmd.Flags().StringP("controller-server-name-override", "", "", "It will override the virtual host name of authority")
	cmd.Flags().StringP("db-connection", "", "/tmp/powerssl.sqlie3", "DB connection")
	cmd.Flags().StringP("db-dialect", "", "sqlite3", "DB Dialect")
	cmd.Flags().StringP("jwks-url", "", "", "JWKS URL")
	cmd.Flags().StringP("metrics-addr", "", ":9090", "HTTP Addr")
	cmd.Flags().StringP("tls-cert-file", "", "", "File containing the default x509 Certificate for GRPC")
	cmd.Flags().StringP("tls-private-key-file", "", "", "File containing the default x509 private key matching --tls-cert-file")
	cmd.Flags().StringP("tracer", "", "jaeger", "Tracing implementation")
	cmd.Flags().StringP("vault-token", "", "", "Vault Token")
	cmd.Flags().StringP("vault-url", "", "", "Vault URL")

	viper.BindPFlag("addr", cmd.Flags().Lookup("addr"))
	viper.BindPFlag("auth-token", cmd.Flags().Lookup("auth-token"))
	viper.BindPFlag("ca-file", cmd.Flags().Lookup("ca-file"))
	viper.BindPFlag("common-name", cmd.Flags().Lookup("common-name"))
	viper.BindPFlag("controller.addr", cmd.Flags().Lookup("controller-addr"))
	viper.BindPFlag("controller.insecure", cmd.Flags().Lookup("controller-insecure"))
	viper.BindPFlag("controller.insecure-skip-tls-verify", cmd.Flags().Lookup("controller-insecure-skip-tls-verify"))
	viper.BindPFlag("controller.server-name-override", cmd.Flags().Lookup("controller-server-name-override"))
	viper.BindPFlag("db.connection", cmd.Flags().Lookup("db-connection"))
	viper.BindPFlag("db.dialect", cmd.Flags().Lookup("db-dialect"))
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

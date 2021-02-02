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
		dbConnection                    string
		dbDialect                       string
		insecure                        bool
		jwksURL                         string
		metricsAddr                     string
		temporalHostPort                string
		temporalNamespace               string
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
			caFile = viper.GetString("ca-file")
			commonName = viper.GetString("common-name")
			dbConnection = viper.GetString("db.connection")
			dbDialect = viper.GetString("db.dialect")
			insecure = viper.GetBool("insecure")
			jwksURL = viper.GetString("jwks-url")
			if !viper.GetBool("no-metrics") {
				metricsAddr = viper.GetString("metrics-addr")
			}
			temporalHostPort = viper.GetString("temporal.host-port")
			temporalNamespace = viper.GetString("temporal.namespace")
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
				TemporalClientConfig: &apiserver.TemporalClientConfig{
					HostPort: temporalHostPort,
					Namespace: temporalNamespace,
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

	cmd.Flags().BoolP("insecure", "", false, "Do not use TLS for the server")
	cmd.Flags().BoolP("no-metrics", "", false, "Do not serve metrics")
	cmd.Flags().BoolP("no-tracing", "", false, "Do not trace")
	cmd.Flags().StringP("addr", "", ":8080", "GRPC Addr")
	cmd.Flags().StringP("ca-file", "", "", "Certificate authority file")
	cmd.Flags().StringP("common-name", "", "", "API Server common name")
	cmd.Flags().StringP("db-connection", "", "", "DB connection")
	cmd.Flags().StringP("db-dialect", "", "", "DB Dialect")
	cmd.Flags().StringP("jwks-url", "", "", "JWKS URL")
	cmd.Flags().StringP("metrics-addr", "", ":9090", "HTTP Addr")
	cmd.Flags().StringP("temporal-host-port", "", "localhost:7233", "Host and port for this client to connect to")
	cmd.Flags().StringP("temporal-namespace", "", "powerssl", "Namespace name for this client to work with")
	cmd.Flags().StringP("tls-cert-file", "", "", "File containing the default x509 Certificate for GRPC")
	cmd.Flags().StringP("tls-private-key-file", "", "", "File containing the default x509 private key matching --tls-cert-file")
	cmd.Flags().StringP("tracer", "", "jaeger", "Tracing implementation")
	cmd.Flags().StringP("vault-token", "", "", "Vault Token")
	cmd.Flags().StringP("vault-url", "", "", "Vault URL")

	viper.BindPFlag("addr", cmd.Flags().Lookup("addr"))
	viper.BindPFlag("ca-file", cmd.Flags().Lookup("ca-file"))
	viper.BindPFlag("common-name", cmd.Flags().Lookup("common-name"))
	viper.BindPFlag("db.connection", cmd.Flags().Lookup("db-connection"))
	viper.BindPFlag("db.dialect", cmd.Flags().Lookup("db-dialect"))
	viper.BindPFlag("insecure", cmd.Flags().Lookup("insecure"))
	viper.BindPFlag("jwks-url", cmd.Flags().Lookup("jwks-url"))
	viper.BindPFlag("metrics-addr", cmd.Flags().Lookup("metrics-addr"))
	viper.BindPFlag("no-metrics", cmd.Flags().Lookup("no-metrics"))
	viper.BindPFlag("no-tracing", cmd.Flags().Lookup("no-tracing"))
	viper.BindPFlag("temporal.host-port", cmd.Flags().Lookup("temporal-host-port"))
	viper.BindPFlag("temporal.namespace", cmd.Flags().Lookup("temporal-namespace"))
	viper.BindPFlag("tls.cert-file", cmd.Flags().Lookup("tls-cert-file"))
	viper.BindPFlag("tls.private-key-file", cmd.Flags().Lookup("tls-private-key-file"))
	viper.BindPFlag("tracer", cmd.Flags().Lookup("tracer"))
	viper.BindPFlag("vault.token", cmd.Flags().Lookup("vault-token"))
	viper.BindPFlag("vault.url", cmd.Flags().Lookup("vault-url"))

	return cmd
}

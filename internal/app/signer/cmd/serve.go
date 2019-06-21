package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"powerssl.io/powerssl/internal/app/signer"
)

func newCmdServe() *cobra.Command {
	var (
		addr              string
		caFile            string
		commonName        string
		insecure          bool
		metricsAddr       string
		tlsCertFile       string
		tlsPrivateKeyFile string
		tracer            string
		vaultToken        string
		vaultURL          string
	)

	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Serve the Signer",
		PreRun: func(cmd *cobra.Command, args []string) {
			addr = viper.GetString("addr")
			caFile = viper.GetString("ca-file")
			commonName = viper.GetString("common-name")
			insecure = viper.GetBool("insecure")
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
			signer.Run(&signer.Config{
				MetricsAddr: metricsAddr,
				ServerConfig: &signer.ServerConfig{
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
			})
		},
	}

	cmd.Flags().BoolP("insecure", "", false, "Do not use TLS for the server")
	cmd.Flags().BoolP("no-metrics", "", false, "Do not serve metrics")
	cmd.Flags().BoolP("no-tracing", "", false, "Do not trace")
	cmd.Flags().StringP("addr", "", ":8080", "GRPC Addr")
	cmd.Flags().StringP("ca-file", "", "", "Certificate authority file")
	cmd.Flags().StringP("common-name", "", "", "API Server common name")
	cmd.Flags().StringP("metrics-addr", "", ":9090", "HTTP Addr")
	cmd.Flags().StringP("tls-cert-file", "", "", "File containing the default x509 Certificate for GRPC.")
	cmd.Flags().StringP("tls-private-key-file", "", "", "File containing the default x509 private key matching --tls-cert-file.")
	cmd.Flags().StringP("tracer", "", "jaeger", "Tracing implementation")
	cmd.Flags().StringP("vault-token", "", "", "Vault Token")
	cmd.Flags().StringP("vault-url", "", "", "Vault URL")

	viper.BindPFlag("addr", cmd.Flags().Lookup("addr"))
	viper.BindPFlag("ca-file", cmd.Flags().Lookup("ca-file"))
	viper.BindPFlag("common-name", cmd.Flags().Lookup("common-name"))
	viper.BindPFlag("insecure", cmd.Flags().Lookup("insecure"))
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

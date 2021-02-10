package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"powerssl.dev/powerssl/internal/app/signer"
	cmdutil "powerssl.dev/powerssl/internal/pkg/cmd"
)

func newCmdServe() *cobra.Command {
	var config signer.Config
	var noMetrics, noTracing bool

	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Serve the Signer",
		Args:  cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.Unmarshal(&config); err != nil {
				return err
			}
			if noMetrics {
				config.Metrics.Addr = ""
			}
			if noTracing {
				config.Tracer = ""
			}
			return config.Validate()
		},
		Run: cmdutil.HandleError(func(cmd *cobra.Command, args []string) error {
			return signer.Run(&config)
		}),
	}

	cmd.Flags().BoolVar(&noMetrics, "no-metrics", false, "Do not serve metrics")
	cmd.Flags().BoolVar(&noTracing, "no-tracing", false, "Do not trace")
	cmd.Flags().Bool("insecure", false, "Do not use TLS for the server")
	cmd.Flags().String("addr", ":8080", "GRPC Addr")
	cmd.Flags().String("ca-file", "", "Certificate authority file")
	cmd.Flags().String("common-name", "", "API Server common name")
	cmd.Flags().String("metrics-addr", ":9090", "HTTP Addr")
	cmd.Flags().String("tls-cert-file", "", "File containing the default x509 Certificate for GRPC.")
	cmd.Flags().String("tls-private-key-file", "", "File containing the default x509 private key matching --tls-cert-file.")
	cmd.Flags().String("tracer", "jaeger", "Tracing implementation")
	cmd.Flags().String("vault-token", "", "Vault Token")
	cmd.Flags().String("vault-url", "", "Vault URL")

	cmdutil.Must(viper.BindPFlag("addr", cmd.Flags().Lookup("addr")))
	cmdutil.Must(viper.BindPFlag("ca-file", cmd.Flags().Lookup("ca-file")))
	cmdutil.Must(viper.BindPFlag("common-name", cmd.Flags().Lookup("common-name")))
	cmdutil.Must(viper.BindPFlag("insecure", cmd.Flags().Lookup("insecure")))
	cmdutil.Must(viper.BindPFlag("metrics.addr", cmd.Flags().Lookup("metrics-addr")))
	cmdutil.Must(viper.BindPFlag("tls-cert-file", cmd.Flags().Lookup("tls-cert-file")))
	cmdutil.Must(viper.BindPFlag("tls-private-key-file", cmd.Flags().Lookup("tls-private-key-file")))
	cmdutil.Must(viper.BindPFlag("tracer", cmd.Flags().Lookup("tracer")))
	cmdutil.Must(viper.BindPFlag("vault-token", cmd.Flags().Lookup("vault-token")))
	cmdutil.Must(viper.BindPFlag("vault-url", cmd.Flags().Lookup("vault-url")))

	return cmd
}

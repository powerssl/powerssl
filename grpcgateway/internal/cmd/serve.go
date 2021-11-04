package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	cmdutil "powerssl.dev/common/cmd"

	"powerssl.dev/grpcgateway/internal"
)

func newCmdServe() *cobra.Command {
	var config *internal.Config
	var noMetrics bool

	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Serve the gRPC Gateway",
		Args:  cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.Unmarshal(&config); err != nil {
				return err
			}
			if noMetrics {
				config.Metrics.Addr = ""
			}
			if !viper.IsSet("apiserver.ca-file") || config.APIServerClient.CAFile == "" {
				config.APIServerClient.CAFile = viper.GetString("ca-file")
			}
			return config.Validate()
		},
		Run: cmdutil.Run(func(ctx context.Context) ([]func() error, func(), error) {
			return internal.Initialize(ctx, config)
		}),
	}

	cmd.Flags().BoolVar(&noMetrics, "no-metrics", false, "Do not serve metrics")
	cmd.Flags().Bool("apiserver-insecure", false, "Use insecure communication")
	cmd.Flags().Bool("apiserver-insecure-skip-tls-verify", false, "Accepts any certificate presented by the server and any host name in that certificate")
	cmd.Flags().String("addr", ":8080", "GRPC Addr")
	cmd.Flags().String("apiserver-addr", "", "GRPC address of API server")
	cmd.Flags().String("apiserver-server-name-override", "", "It will override the virtual host name of authority")
	cmd.Flags().String("ca-file", "", "Certificate authority file")
	cmd.Flags().String("metrics-addr", ":9090", "HTTP Addr")

	cmdutil.Must(viper.BindPFlag("addr", cmd.Flags().Lookup("addr")))
	cmdutil.Must(viper.BindPFlag("apiserver.addr", cmd.Flags().Lookup("apiserver-addr")))
	cmdutil.Must(viper.BindPFlag("apiserver.ca-file", cmd.Flags().Lookup("ca-file")))
	cmdutil.Must(viper.BindPFlag("apiserver.insecure", cmd.Flags().Lookup("apiserver-insecure")))
	cmdutil.Must(viper.BindPFlag("apiserver.insecure-skip-tls-verify", cmd.Flags().Lookup("apiserver-insecure-skip-tls-verify")))
	cmdutil.Must(viper.BindPFlag("apiserver.server-name-override", cmd.Flags().Lookup("apiserver-server-name-override")))
	cmdutil.Must(viper.BindPFlag("ca-file", cmd.Flags().Lookup("ca-file")))
	cmdutil.Must(viper.BindPFlag("metrics.addr", cmd.Flags().Lookup("metrics-addr")))

	return cmd
}

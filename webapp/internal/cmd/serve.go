package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	cmdutil "powerssl.dev/common/cmd"
	"powerssl.dev/webapp/internal"
)

func newCmdServe() *cobra.Command {
	var config internal.Config
	var noMetrics bool

	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Serve the WebApp",
		Args:  cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.Unmarshal(&config); err != nil {
				return err
			}
			if noMetrics {
				config.Metrics.Addr = ""
			}
			return config.Validate()
		},
		Run: cmdutil.HandleError(func(cmd *cobra.Command, args []string) error {
			return internal.Run(&config)
		}),
	}

	cmd.Flags().BoolVar(&noMetrics, "no-metrics", false, "Do not serve metrics")
	cmd.Flags().Bool("insecure", false, "Do not use TLS for the server")
	cmd.Flags().StringP("addr", "", ":8080", "Addr")
	cmd.Flags().String("apiserver-addr", "", "GRPC address of API server")
	cmd.Flags().StringP("auth-uri", "", "", "Auth URI")
	cmd.Flags().StringP("grpcweb-uri", "", "", "gRPC-Web URI")
	cmd.Flags().StringP("metrics-addr", "", ":9090", "HTTP Addr")

	cmdutil.Must(viper.BindPFlag("addr", cmd.Flags().Lookup("addr")))
	cmdutil.Must(viper.BindPFlag("apiserver.addr", cmd.Flags().Lookup("apiserver-addr")))
	cmdutil.Must(viper.BindPFlag("auth.uri", cmd.Flags().Lookup("auth-uri")))
	cmdutil.Must(viper.BindPFlag("grpcweb.uri", cmd.Flags().Lookup("grpcweb-uri")))
	cmdutil.Must(viper.BindPFlag("metrics.addr", cmd.Flags().Lookup("metrics-addr")))
	cmdutil.Must(viper.BindPFlag("insecure", cmd.Flags().Lookup("insecure")))

	return cmd
}

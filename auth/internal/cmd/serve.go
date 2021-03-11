package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"powerssl.dev/auth/internal"
	cmdutil "powerssl.dev/common/cmd"
)

func newCmdServe() *cobra.Command {
	var config internal.Config
	var noMetrics bool

	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Serve the API",
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
	cmd.Flags().String("addr", ":8080", "GRPC")
	cmd.Flags().String("jwt-private-key-file", "", "JWT private key file")
	cmd.Flags().String("metrics-addr", ":9090", "HTTP Addr")
	cmd.Flags().String("webapp-uri", "", "WebApp URI")

	cmdutil.Must(viper.BindPFlag("addr", cmd.Flags().Lookup("addr")))
	cmdutil.Must(viper.BindPFlag("jwt.private-key-file", cmd.Flags().Lookup("jwt-private-key-file")))
	cmdutil.Must(viper.BindPFlag("metrics.addr", cmd.Flags().Lookup("metrics-addr")))
	cmdutil.Must(viper.BindPFlag("insecure", cmd.Flags().Lookup("insecure")))
	cmdutil.Must(viper.BindPFlag("webapp.uri", cmd.Flags().Lookup("webapp-uri")))

	return cmd
}

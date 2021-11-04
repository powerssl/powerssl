package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	cmdutil "powerssl.dev/common/cmd"

	"powerssl.dev/auth/internal"
)

func newCmdServe() *cobra.Command {
	var config *internal.Config
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
		Run: cmdutil.Run(func(ctx context.Context) ([]func() error, func(), error) {
			return internal.Initialize(ctx, config)
		}),
	}

	// TODO: Config doesn't map anymore
	cmd.Flags().BoolVar(&noMetrics, "no-metrics", false, "Do not serve metrics")
	cmd.Flags().Bool("insecure", false, "Do not use TLS for the server")
	cmd.Flags().String("addr", ":8080", "GRPC")
	cmd.Flags().String("auth-uri", "", "Auth URI")
	cmd.Flags().String("jwt-private-key-file", "", "JWT private key file")
	cmd.Flags().String("metrics-addr", ":9090", "HTTP Addr")
	cmd.Flags().String("oauth2-github-client-id", "", "Oauth2 GitHub ClientID")
	cmd.Flags().String("oauth2-github-client-secret", "", "Oauth2 GitHub ClientSecret")
	cmd.Flags().String("webapp-uri", "", "WebApp URI")
	// TODO: TLS

	cmdutil.Must(viper.BindPFlag("addr", cmd.Flags().Lookup("addr")))
	cmdutil.Must(viper.BindPFlag("auth.uri", cmd.Flags().Lookup("auth-uri")))
	cmdutil.Must(viper.BindPFlag("jwt.private-key-file", cmd.Flags().Lookup("jwt-private-key-file")))
	cmdutil.Must(viper.BindPFlag("metrics.addr", cmd.Flags().Lookup("metrics-addr")))
	cmdutil.Must(viper.BindPFlag("insecure", cmd.Flags().Lookup("insecure")))
	cmdutil.Must(viper.BindPFlag("oauth2.github.client-id", cmd.Flags().Lookup("oauth2-github-client-id")))
	cmdutil.Must(viper.BindPFlag("oauth2.github.client-secret", cmd.Flags().Lookup("oauth2-github-client-secret")))
	cmdutil.Must(viper.BindPFlag("webapp.uri", cmd.Flags().Lookup("webapp-uri")))

	return cmd
}

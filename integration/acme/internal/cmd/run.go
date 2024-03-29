package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	cmdutil "powerssl.dev/common/cmd"
	"powerssl.dev/sdk/integration"
	"powerssl.dev/sdk/integration/vault"

	"powerssl.dev/integration/acme/internal"
)

func newCmdRun() *cobra.Command {
	var config integration.Config
	var noMetrics, noTracing bool

	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run ACME integration",
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
			if !viper.IsSet("controller.ca-file") || config.ControllerClientConfig.CAFile == "" {
				config.ControllerClientConfig.CAFile = viper.GetString("ca-file")
			}
			return config.Validate()
		},
		Run: cmdutil.HandleError(func(cmd *cobra.Command, args []string) error {
			handler, err := acme.New(vault.Config{
				Address: "https://localhost:8200",
				CAFile:  config.ControllerClientConfig.CAFile, // TODO: Wrong config
			})
			if err != nil {
				return err
			}
			return integration.Run(&config, integration.KindACME, "acme", handler)
		}),
	}

	cmd.Flags().BoolVar(&noMetrics, "no-metrics", false, "Do not serve metrics")
	cmd.Flags().BoolVar(&noTracing, "no-tracing", false, "Do not trace")
	cmd.Flags().Bool("controller-insecure", false, "Use insecure communication")
	cmd.Flags().Bool("controller-insecure-skip-tls-verify", false, "Accepts any certificate presented by the server and any host name in that certificate")
	cmd.Flags().String("auth-token", "", "Authentication token")
	cmd.Flags().String("ca-file", "", "Certificate authority file")
	cmd.Flags().String("controller-addr", "", "GRPC address of Controller")
	cmd.Flags().String("controller-server-name-override", "", "It will override the virtual host name of authority")
	cmd.Flags().String("metrics-addr", ":9090", "HTTP Addr")
	cmd.Flags().String("tracer", "jaeger", "Tracing implementation")

	cmdutil.Must(viper.BindPFlag("auth-token", cmd.Flags().Lookup("auth-token")))
	cmdutil.Must(viper.BindPFlag("ca-file", cmd.Flags().Lookup("ca-file")))
	cmdutil.Must(viper.BindPFlag("controller.addr", cmd.Flags().Lookup("controller-addr")))
	cmdutil.Must(viper.BindPFlag("controller.ca-file", cmd.Flags().Lookup("ca-file")))
	cmdutil.Must(viper.BindPFlag("controller.insecure", cmd.Flags().Lookup("controller-insecure")))
	cmdutil.Must(viper.BindPFlag("controller.insecure-skip-tls-verify", cmd.Flags().Lookup("controller-insecure-skip-tls-verify")))
	cmdutil.Must(viper.BindPFlag("controller.server-name-override", cmd.Flags().Lookup("controller-server-name-override")))
	cmdutil.Must(viper.BindPFlag("metrics.addr", cmd.Flags().Lookup("metrics-addr")))
	cmdutil.Must(viper.BindPFlag("tracer", cmd.Flags().Lookup("tracer")))

	return cmd
}

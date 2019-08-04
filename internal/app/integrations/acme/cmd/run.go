package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"powerssl.dev/powerssl/internal/app/integrations/acme"
	"powerssl.dev/powerssl/pkg/integration"
)

func newCmdRun() *cobra.Command {
	var (
		addr                  string
		authToken             string
		caFile                string
		insecure              bool
		insecureSkipTLSVerify bool
		metricsAddr           string
		serverNameOverride    string
		tracer                string
	)

	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run ACME integration",
		PreRun: func(cmd *cobra.Command, args []string) {
			addr = viper.GetString("addr")
			authToken = viper.GetString("auth-token")
			caFile = viper.GetString("ca-file")
			insecure = viper.GetBool("insecure")
			insecureSkipTLSVerify = viper.GetBool("insecure-skip-tls-verify")
			if !viper.GetBool("no-metrics") {
				metricsAddr = viper.GetString("metrics-addr")
			}
			serverNameOverride = viper.GetString("server-name-override")
			if !viper.GetBool("no-tracing") {
				tracer = viper.GetString("tracer")
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			integration.Run(&integration.Config{
				AuthToken: authToken,
				ControllerClientConfig: &integration.ControllerClientConfig{
					Addr:                  addr,
					CAFile:                caFile,
					Insecure:              insecure,
					InsecureSkipTLSVerify: insecureSkipTLSVerify,
					ServerNameOverride:    serverNameOverride,
				},
				MetricsAddr: metricsAddr,
				Tracer:      tracer,
			}, integration.KindACME, "acme", acme.New())
		},
	}

	cmd.Flags().BoolP("insecure", "", false, "Use insecure communication")
	cmd.Flags().BoolP("insecure-skip-tls-verify", "", false, "Accepts any certificate presented by the server and any host name in that certificate")
	cmd.Flags().BoolP("no-metrics", "", false, "Do not serve metrics")
	cmd.Flags().BoolP("no-tracing", "", false, "Do not trace")
	cmd.Flags().StringP("addr", "", "", "GRPC address of Controller")
	cmd.Flags().StringP("auth-token", "", "", "Authentication token")
	cmd.Flags().StringP("ca-file", "", "", "Certificate authority file")
	cmd.Flags().StringP("metrics-addr", "", ":9090", "HTTP Addr")
	cmd.Flags().StringP("server-name-override", "", "", "It will override the virtual host name of authority")
	cmd.Flags().StringP("tracer", "", "jaeger", "Tracing implementation")

	viper.BindPFlag("addr", cmd.Flags().Lookup("addr"))
	viper.BindPFlag("auth-token", cmd.Flags().Lookup("auth-token"))
	viper.BindPFlag("ca-file", cmd.Flags().Lookup("ca-file"))
	viper.BindPFlag("insecure", cmd.Flags().Lookup("insecure"))
	viper.BindPFlag("insecure-skip-tls-verify", cmd.Flags().Lookup("insecure-skip-tls-verify"))
	viper.BindPFlag("metrics-addr", cmd.Flags().Lookup("metrics-addr"))
	viper.BindPFlag("no-metrics", cmd.Flags().Lookup("no-metrics"))
	viper.BindPFlag("no-tracing", cmd.Flags().Lookup("no-tracing"))
	viper.BindPFlag("server-name-override", cmd.Flags().Lookup("server-name-override"))
	viper.BindPFlag("tracer", cmd.Flags().Lookup("tracer"))

	return cmd
}

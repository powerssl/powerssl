package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"powerssl.io/powerssl/internal/app/grpcgateway"
)

func newCmdServe() *cobra.Command {
	var (
		addr                           string
		apiserverAddr                  string
		apiserverInsecure              bool
		apiserverInsecureSkipTLSVerify bool
		apiserverServerNameOverride    string
		caFile                         string
		metricsAddr                    string
	)

	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Serve the gRPC Gateway",
		PreRun: func(cmd *cobra.Command, args []string) {
			addr = viper.GetString("addr")
			apiserverAddr = viper.GetString("apiserver.addr")
			apiserverInsecure = viper.GetBool("apiserver.insecure")
			apiserverInsecureSkipTLSVerify = viper.GetBool("apiserver.insecure-skip-tls-verify")
			apiserverServerNameOverride = viper.GetString("apiserver.server-name-override")
			caFile = viper.GetString("ca-file")
			if !viper.GetBool("no-metrics") {
				metricsAddr = viper.GetString("metrics-addr")
			}

		},
		Run: func(cmd *cobra.Command, args []string) {
			grpcgateway.Run(&grpcgateway.Config{
				APIServerClientConfig: &grpcgateway.APIServerClientConfig{
					Addr:                  apiserverAddr,
					CAFile:                caFile,
					Insecure:              apiserverInsecure,
					InsecureSkipTLSVerify: apiserverInsecureSkipTLSVerify,
					ServerNameOverride:    apiserverServerNameOverride,
				},
				Addr:        addr,
				MetricsAddr: metricsAddr,
			})
		},
	}

	cmd.Flags().BoolP("apiserver-insecure", "", false, "Use insecure communication")
	cmd.Flags().BoolP("apiserver-insecure-skip-tls-verify", "", false, "Accepts any certificate presented by the server and any host name in that certificate")
	cmd.Flags().BoolP("no-metrics", "", false, "Do not serve metrics")
	cmd.Flags().StringP("addr", "", ":8080", "Addr")
	cmd.Flags().StringP("apiserver-addr", "", "", "GRPC address of APIServer")
	cmd.Flags().StringP("apiserver-auth-token", "", "", "APIServer authentication token")
	cmd.Flags().StringP("apiserver-server-name-override", "", "", "It will override the virtual host name of authority")
	cmd.Flags().StringP("ca-file", "", "", "Certificate authority file")
	cmd.Flags().StringP("metrics-addr", "", ":9090", "HTTP Addr")

	viper.BindPFlag("addr", cmd.Flags().Lookup("addr"))
	viper.BindPFlag("apiserver.addr", cmd.Flags().Lookup("apiserver-addr"))
	viper.BindPFlag("apiserver.auth-token", cmd.Flags().Lookup("apiserver-auth-token"))
	viper.BindPFlag("apiserver.insecure", cmd.Flags().Lookup("apiserver-insecure"))
	viper.BindPFlag("apiserver.insecure-skip-tls-verify", cmd.Flags().Lookup("apiserver-insecure-skip-tls-verify"))
	viper.BindPFlag("apiserver.server-name-override", cmd.Flags().Lookup("apiserver-server-name-override"))
	viper.BindPFlag("ca-file", cmd.Flags().Lookup("ca-file"))
	viper.BindPFlag("metrics-addr", cmd.Flags().Lookup("metrics-addr"))
	viper.BindPFlag("no-metrics", cmd.Flags().Lookup("no-metrics"))

	return cmd
}

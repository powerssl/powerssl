package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"powerssl.dev/powerssl/internal/app/worker"
)

func newCmdRun() *cobra.Command {
	var (
		apiserverAddr                  string
		apiserverInsecure              bool
		apiserverInsecureSkipTLSVerify bool
		apiserverServerNameOverride    string
		authToken                      string
		caFile                         string
		metricsAddr                    string
		temporalHostPort               string
		temporalNamespace              string
		tracer                         string
		vaultToken                     string
		vaultURL                       string
	)

	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run the Worker",
		PreRun: func(cmd *cobra.Command, args []string) {
			apiserverAddr = viper.GetString("apiserver.addr")
			apiserverInsecure = viper.GetBool("apiserver.insecure")
			apiserverInsecureSkipTLSVerify = viper.GetBool("apiserver.insecure-skip-tls-verify")
			apiserverServerNameOverride = viper.GetString("apiserver.server-name-override")
			authToken = viper.GetString("auth-token")
			caFile = viper.GetString("ca-file")
			if !viper.GetBool("no-metrics") {
				metricsAddr = viper.GetString("metrics-addr")
			}
			temporalHostPort = viper.GetString("temporal.host-port")
			temporalNamespace = viper.GetString("temporal.namespace")
			if !viper.GetBool("no-tracing") {
				tracer = viper.GetString("tracer")
			}
			vaultToken = viper.GetString("vault.token")
			vaultURL = viper.GetString("vault.url")
		},
		Run: func(cmd *cobra.Command, args []string) {
			worker.Run(&worker.Config{
				AuthToken: authToken,
				APIServerClientConfig: &worker.APIServerClientConfig{
					CAFile:                caFile,
					Addr:                  apiserverAddr,
					Insecure:              apiserverInsecure,
					InsecureSkipTLSVerify: apiserverInsecureSkipTLSVerify,
					ServerNameOverride:    apiserverServerNameOverride,
				},
				MetricsAddr: metricsAddr,
				TemporalClientConfig: &worker.TemporalClientConfig{
					HostPort: temporalHostPort,
					Namespace: temporalNamespace,
				},
				Tracer: tracer,
				VaultClientConfig: &worker.VaultClientConfig{
					CAFile: caFile,
					Token:  vaultToken,
					URL:    vaultURL,
				},
			})
		},
	}

	cmd.Flags().BoolP("apiserver-insecure", "", false, "Use insecure communication")
	cmd.Flags().BoolP("apiserver-insecure-skip-tls-verify", "", false, "Accepts any certificate presented by the server and any host name in that certificate")
	cmd.Flags().BoolP("no-metrics", "", false, "Do not serve metrics")
	cmd.Flags().BoolP("no-tracing", "", false, "Do not trace")
	cmd.Flags().StringP("addr", "", ":8080", "GRPC Addr")
	cmd.Flags().StringP("apiserver-addr", "", "", "GRPC address of API server")
	cmd.Flags().StringP("apiserver-server-name-override", "", "", "It will override the virtual host name of authority")
	cmd.Flags().StringP("auth-token", "", "", "Authentication token")
	cmd.Flags().StringP("ca-file", "", "", "Certificate authority file")
	cmd.Flags().StringP("metrics-addr", "", ":9090", "HTTP Addr")
	cmd.Flags().StringP("temporal-host-port", "", "localhost:7233", "Host and port for this client to connect to")
	cmd.Flags().StringP("temporal-namespace", "", "powerssl", "Namespace name for this client to work with")
	cmd.Flags().StringP("tracer", "", "jaeger", "Tracing implementation")
	cmd.Flags().StringP("vault-token", "", "", "Vault Token")
	cmd.Flags().StringP("vault-url", "", "", "Vault URL")

	viper.BindPFlag("addr", cmd.Flags().Lookup("addr"))
	viper.BindPFlag("apiserver.addr", cmd.Flags().Lookup("apiserver-addr"))
	viper.BindPFlag("apiserver.insecure", cmd.Flags().Lookup("apiserver-insecure"))
	viper.BindPFlag("apiserver.insecure-skip-tls-verify", cmd.Flags().Lookup("apiserver-insecure-skip-tls-verify"))
	viper.BindPFlag("apiserver.server-name-override", cmd.Flags().Lookup("apiserver-server-name-override"))
	viper.BindPFlag("auth-token", cmd.Flags().Lookup("auth-token"))
	viper.BindPFlag("ca-file", cmd.Flags().Lookup("ca-file"))
	viper.BindPFlag("metrics-addr", cmd.Flags().Lookup("metrics-addr"))
	viper.BindPFlag("no-metrics", cmd.Flags().Lookup("no-metrics"))
	viper.BindPFlag("no-tracing", cmd.Flags().Lookup("no-tracing"))
	viper.BindPFlag("temporal.host-port", cmd.Flags().Lookup("temporal-host-port"))
	viper.BindPFlag("temporal.namespace", cmd.Flags().Lookup("temporal-namespace"))
	viper.BindPFlag("tracer", cmd.Flags().Lookup("tracer"))
	viper.BindPFlag("vault.token", cmd.Flags().Lookup("vault-token"))
	viper.BindPFlag("vault.url", cmd.Flags().Lookup("vault-url"))

	return cmd
}

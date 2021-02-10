package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"powerssl.dev/powerssl/internal/app/worker"
)

func newCmdRun() *cobra.Command {
	var config worker.Config

	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run the Worker",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			//config.AuthToken = viper.GetString("auth-token")
			//config.APIServerClientConfig.CAFile = viper.GetString("ca-file")
			//config.APIServerClientConfig.Addr = viper.GetString("apiserver.addr")
			//config.APIServerClientConfig.Insecure = viper.GetBool("apiserver.insecure")
			//config.APIServerClientConfig.InsecureSkipTLSVerify = viper.GetBool("apiserver.insecure-skip-tls-verify")
			//config.APIServerClientConfig.ServerNameOverride = viper.GetString("apiserver.server-name-override")
			//if !viper.GetBool("no-metrics") {
			//	config.MetricsAddr = viper.GetString("metrics-addr")
			//}
			//config.TemporalClientConfig.CAFile = viper.GetString("ca-file")
			//config.TemporalClientConfig.HostPort = viper.GetString("temporal.host-port")
			//config.TemporalClientConfig.Namespace = viper.GetString("temporal.namespace")
			//if !viper.GetBool("no-tracing") {
			//	config.Tracer = viper.GetString("tracer")
			//}
			//config.VaultClientConfig.CAFile = viper.GetString("ca-file")
			//config.VaultClientConfig.Token = viper.GetString("vault.token")
			//config.VaultClientConfig.URL = viper.GetString("vault.url")
			return viper.Unmarshal(&config)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(fmt.Sprintf("%+v", config))
			return worker.Run(&config)
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

	must(viper.BindPFlag("addr", cmd.Flags().Lookup("addr")))
	must(viper.BindPFlag("apiserver.addr", cmd.Flags().Lookup("apiserver-addr")))
	must(viper.BindPFlag("apiserver.insecure", cmd.Flags().Lookup("apiserver-insecure")))
	must(viper.BindPFlag("apiserver.insecure-skip-tls-verify", cmd.Flags().Lookup("apiserver-insecure-skip-tls-verify")))
	must(viper.BindPFlag("apiserver.server-name-override", cmd.Flags().Lookup("apiserver-server-name-override")))
	must(viper.BindPFlag("auth-token", cmd.Flags().Lookup("auth-token")))
	must(viper.BindPFlag("ca-file", cmd.Flags().Lookup("ca-file")))
	must(viper.BindPFlag("metrics-addr", cmd.Flags().Lookup("metrics-addr")))
	must(viper.BindPFlag("no-metrics", cmd.Flags().Lookup("no-metrics")))
	must(viper.BindPFlag("no-tracing", cmd.Flags().Lookup("no-tracing")))
	must(viper.BindPFlag("temporal.host-port", cmd.Flags().Lookup("temporal-host-port")))
	must(viper.BindPFlag("temporal.namespace", cmd.Flags().Lookup("temporal-namespace")))
	must(viper.BindPFlag("tracer", cmd.Flags().Lookup("tracer")))
	must(viper.BindPFlag("vault.token", cmd.Flags().Lookup("vault-token")))
	must(viper.BindPFlag("vault.url", cmd.Flags().Lookup("vault-url")))

	return cmd
}

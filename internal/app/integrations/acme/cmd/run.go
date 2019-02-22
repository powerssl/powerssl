package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"powerssl.io/internal/app/integrations/acme"
	"powerssl.io/pkg/integration"
)

func newCmdRun() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run ACME integration",
		Run: func(cmd *cobra.Command, args []string) {
			addr := viper.GetString("addr")
			certFile := viper.GetString("ca-file")
			insecure := viper.GetBool("insecure")
			insecureSkipTLSVerify := viper.GetBool("insecure-skip-tls-verify")
			var metricsAddr string
			if !viper.GetBool("no-metrics") {
				metricsAddr = viper.GetString("metrics-addr")
			}
			serverNameOverride := viper.GetString("server-name-override")
			var tracer string
			if !viper.GetBool("no-tracing") {
				tracer = viper.GetString("tracer")
			}

			ok := true
			if addr == "" {
				ok = false
				fmt.Println("Provide addr")
			}
			if !insecure && !insecureSkipTLSVerify && certFile == "" {
				ok = false
				fmt.Println("Provide ca-file")
			}
			if !ok {
				os.Exit(1)
			}

			integration.Run(addr, certFile, serverNameOverride, insecure, insecureSkipTLSVerify, metricsAddr, tracer, integration.KindACME, "acme", acme.New())
		},
	}

	cmd.Flags().BoolP("insecure", "", false, "Use insecure communication")
	cmd.Flags().BoolP("insecure-skip-tls-verify", "", false, "Accepts any certificate presented by the server and any host name in that certificate")
	cmd.Flags().BoolP("no-metrics", "", false, "Do not serve metrics")
	cmd.Flags().BoolP("no-tracing", "", false, "Do not trace")
	cmd.Flags().StringP("addr", "", "", "GRPC address of Controller")
	cmd.Flags().StringP("ca-file", "", "", "Certificate authority file")
	cmd.Flags().StringP("metrics-addr", "", ":9090", "HTTP Addr")
	cmd.Flags().StringP("server-name-override", "", "", "It will override the virtual host name of authority")
	cmd.Flags().StringP("tracer", "", "jaeger", "Tracing implementation")

	viper.BindPFlag("addr", cmd.Flags().Lookup("addr"))
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

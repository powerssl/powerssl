package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"powerssl.io/pkg/integration"
	"powerssl.io/pkg/integrations/cloudflare"
)

var runCmd = &cobra.Command{
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

		integ := integration.New(addr, certFile, serverNameOverride, insecure, insecureSkipTLSVerify, metricsAddr, tracer, integration.KindDNS, "cloudflare", cloudflare.New())
		integ.Run()
	},
}

func init() {
	runCmd.Flags().BoolP("insecure", "", false, "Use insecure communication")
	runCmd.Flags().BoolP("insecure-skip-tls-verify", "", false, "Accepts any certificate presented by the server and any host name in that certificate")
	runCmd.Flags().BoolP("no-metrics", "", false, "Do not serve metrics")
	runCmd.Flags().BoolP("no-tracing", "", false, "Do not trace")
	runCmd.Flags().StringP("addr", "", "", "GRPC address of Controller")
	runCmd.Flags().StringP("ca-file", "", "", "Certificate authority file")
	runCmd.Flags().StringP("metrics-addr", "", ":9090", "HTTP Addr")
	runCmd.Flags().StringP("server-name-override", "", "", "It will override the virtual host name of authority")
	runCmd.Flags().StringP("tracer", "", "jaeger", "Tracing implementation")

	viper.BindPFlag("addr", runCmd.Flags().Lookup("addr"))
	viper.BindPFlag("ca-file", runCmd.Flags().Lookup("ca-file"))
	viper.BindPFlag("insecure", runCmd.Flags().Lookup("insecure"))
	viper.BindPFlag("insecure-skip-tls-verify", runCmd.Flags().Lookup("insecure-skip-tls-verify"))
	viper.BindPFlag("metrics-addr", runCmd.Flags().Lookup("metrics-addr"))
	viper.BindPFlag("no-metrics", runCmd.Flags().Lookup("no-metrics"))
	viper.BindPFlag("no-tracing", runCmd.Flags().Lookup("no-tracing"))
	viper.BindPFlag("server-name-override", runCmd.Flags().Lookup("server-name-override"))
	viper.BindPFlag("tracer", runCmd.Flags().Lookup("tracer"))

	rootCmd.AddCommand(runCmd)
}

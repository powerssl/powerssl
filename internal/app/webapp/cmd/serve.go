package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"powerssl.io/powerssl/internal/app/webapp"
)

func newCmdServe() *cobra.Command {
	var (
		addr        string
		apiAddr     string
		authURI     string
		metricsAddr string
	)

	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Serve the API",
		PreRun: func(cmd *cobra.Command, args []string) {
			addr = viper.GetString("addr")
			apiAddr = viper.GetString("api-addr")
			authURI = viper.GetString("auth-uri")
			if !viper.GetBool("no-metrics") {
				metricsAddr = viper.GetString("metrics-addr")
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			webapp.Run(&webapp.Config{
				Addr:        addr,
				APIAddr:     apiAddr,
				AuthURI:     authURI,
				MetricsAddr: metricsAddr,
			})
		},
	}

	cmd.Flags().BoolP("no-metrics", "", false, "Do not serve metrics")
	cmd.Flags().StringP("addr", "", ":8080", "Addr")
	cmd.Flags().StringP("api-addr", "", "", "API Addr")
	cmd.Flags().StringP("auth-uri", "", "", "Auth URI")
	cmd.Flags().StringP("metrics-addr", "", ":9090", "HTTP Addr")

	viper.BindPFlag("addr", cmd.Flags().Lookup("addr"))
	viper.BindPFlag("api-addr", cmd.Flags().Lookup("api-addr"))
	viper.BindPFlag("auth-uri", cmd.Flags().Lookup("auth-uri"))
	viper.BindPFlag("metrics-addr", cmd.Flags().Lookup("metrics-addr"))
	viper.BindPFlag("no-metrics", cmd.Flags().Lookup("no-metrics"))

	return cmd
}

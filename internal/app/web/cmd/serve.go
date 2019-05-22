package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"powerssl.io/powerssl/internal/app/web"
)

func newCmdServe() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Serve the API",
		Run: func(cmd *cobra.Command, args []string) {
			addr := viper.GetString("addr")
			authURI := viper.GetString("auth-uri")
			apiAddr := viper.GetString("api-addr")
			var metricsAddr string
			if !viper.GetBool("no-metrics") {
				metricsAddr = viper.GetString("metrics-addr")
			}

			ok := true
			if addr == "" {
				ok = false
				fmt.Println("Provide addr")
			}
			if authURI == "" {
				ok = false
				fmt.Println("Provide auth-uri")
			}
			if apiAddr == "" {
				ok = false
				fmt.Println("Provide api-addr")
			}
			if !ok {
				os.Exit(1)
			}

			web.Run(addr, metricsAddr, authURI, apiAddr)
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

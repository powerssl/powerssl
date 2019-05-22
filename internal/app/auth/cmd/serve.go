package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"powerssl.io/powerssl/internal/app/auth"
)

func newCmdServe() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Serve the API",
		Run: func(cmd *cobra.Command, args []string) {
			addr := viper.GetString("addr")
			var metricsAddr string
			if !viper.GetBool("no-metrics") {
				metricsAddr = viper.GetString("metrics-addr")
			}
			jwtPrivateKeyFile := viper.GetString("jwt.private-key-file")
			webappURI := viper.GetString("webapp-uri")

			ok := true
			if addr == "" {
				ok = false
				fmt.Println("Provide addr")
			}
			if jwtPrivateKeyFile == "" {
				ok = false
				fmt.Println("Provide jwt-private-key-file")
			}
			if webappURI == "" {
				ok = false
				fmt.Println("Provide webapp-uri")
			}
			if !ok {
				os.Exit(1)
			}

			auth.Run(addr, metricsAddr, jwtPrivateKeyFile, webappURI)
		},
	}

	cmd.Flags().BoolP("no-metrics", "", false, "Do not serve metrics")
	cmd.Flags().StringP("addr", "", ":8080", "GRPC")
	cmd.Flags().StringP("jwt-private-key-file", "", "", "JWT private key file")
	cmd.Flags().StringP("metrics-addr", "", ":9090", "HTTP Addr")
	cmd.Flags().StringP("webapp-uri", "", "", "WebApp URI")

	viper.BindPFlag("addr", cmd.Flags().Lookup("addr"))
	viper.BindPFlag("jwt.private-key-file", cmd.Flags().Lookup("jwt-private-key-file"))
	viper.BindPFlag("metrics-addr", cmd.Flags().Lookup("metrics-addr"))
	viper.BindPFlag("no-metrics", cmd.Flags().Lookup("no-metrics"))
	viper.BindPFlag("webapp-uri", cmd.Flags().Lookup("webapp-uri"))

	return cmd
}

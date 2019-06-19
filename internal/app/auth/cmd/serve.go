package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"powerssl.io/powerssl/internal/app/auth"
)

func newCmdServe() *cobra.Command {
	var (
		addr              string
		jwtPrivateKeyFile string
		metricsAddr       string
		webAppURI         string
	)

	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Serve the API",
		PreRun: func(cmd *cobra.Command, args []string) {
			addr = viper.GetString("addr")
			jwtPrivateKeyFile = viper.GetString("jwt.private-key-file")
			if !viper.GetBool("no-metrics") {
				metricsAddr = viper.GetString("metrics-addr")
			}
			webAppURI = viper.GetString("webapp-uri")
		},
		Run: func(cmd *cobra.Command, args []string) {
			auth.Run(&auth.Config{
				Addr:              addr,
				JWTPrivateKeyFile: jwtPrivateKeyFile,
				MetricsAddr:       metricsAddr,
				WebAppURI:         webAppURI,
			})
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

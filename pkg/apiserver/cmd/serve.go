package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"powerssl.io/pkg/apiserver"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the API",
	Run: func(cmd *cobra.Command, args []string) {
		grpcAddr := viper.GetString("grpcAddr")
		apiserver.Run(grpcAddr)
	},
}

func init() {
	serveCmd.Flags().StringP("grpc-addr", "", ":8080", "GRPC Addr")

	viper.BindPFlag("grpcAddr", serveCmd.Flags().Lookup("grpc-addr"))

	rootCmd.AddCommand(serveCmd)
}

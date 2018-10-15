package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"powerssl.io/pkg/controller"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the API",
	Run: func(cmd *cobra.Command, args []string) {
		grpcAddr := viper.GetString("grpc.addr")

		controller.Run(grpcAddr)
	},
}

func init() {
	serveCmd.Flags().StringP("grpc-addr", "", ":8081", "GRPC Addr")

	viper.BindPFlag("grpc.addr", serveCmd.Flags().Lookup("grpc-addr"))

	rootCmd.AddCommand(serveCmd)
}

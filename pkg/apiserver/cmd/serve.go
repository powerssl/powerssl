package cmd

import (
	"github.com/spf13/cobra"

	"powerssl.io/pkg/apiserver"
)

// TODO
const grpcAddr = "localhost:8080"

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the API",
	Run: func(cmd *cobra.Command, args []string) {
		apiserver.Run(grpcAddr)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

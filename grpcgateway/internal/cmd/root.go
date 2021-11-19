package cmd

import (
	"github.com/spangenberg/snakecharmer"
	"github.com/spf13/cobra"

	"powerssl.dev/common/version"
)

func Execute() {
	snakecharmer.ExecuteWithConfig(NewCmdRoot(), "/etc/powerssl/grpcgateway", "powerssl")
}

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "powerssl-grpcgateway",
		Short: "powerssl-grpcgateway provides PowerSSL gRPC Gateway",
		Long: `powerssl-grpcgateway provides PowerSSL gRPC Gareway.

Find more information at: https://docs.powerssl.io/powerssl-grpcgateway`,
		Version: version.String(),
	}

	cmd.AddCommand(newCmdServe())

	return cmd
}

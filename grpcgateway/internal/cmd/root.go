package cmd

import (
	"github.com/spangenberg/snakecharmer"
	"github.com/spf13/cobra"

	"powerssl.dev/common/version"
)

const component = "grpcgateway"

var (
	cfgFile string
	verbose bool
)

func Execute() {
	snakecharmer.ExecuteWithConfig(NewCmdRoot(), "/etc/powerssl/"+component, "powerssl", &cfgFile, &verbose)
}

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "powerssl-grpcgateway",
		Short: "powerssl-grpcgateway provides PowerSSL gRPC Gateway",
		Long: `powerssl-grpcgateway provides PowerSSL gRPC Gareway.

Find more information at: https://docs.powerssl.io/powerssl-grpcgateway`,
		Version: version.String(),
	}

	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is /etc/powerssl/grpcgateway/config.yaml)")

	cmd.AddCommand(newCmdServe())

	return cmd
}

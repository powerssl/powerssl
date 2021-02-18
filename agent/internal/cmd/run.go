package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"powerssl.dev/agent/internal"
	cmdutil "powerssl.dev/common/cmd"
)

func newCmdRun() *cobra.Command {
	var config internal.Config

	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run the Agent",
		Args:  cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := viper.Unmarshal(&config); err != nil {
				return err
			}
			return config.Validate()
		},
		Run: cmdutil.HandleError(func(cmd *cobra.Command, args []string) error {
			return internal.Run(&config)
		}),
	}

	cmd.Flags().Bool("apiserver-insecure", false, "Use insecure communication")
	cmd.Flags().Bool("apiserver-insecure-skip-tls-verify", false, "Accepts any certificate presented by the server and any host name in that certificate")
	cmd.Flags().String("apiserver-addr", "", "GRPC address of API server")
	cmd.Flags().String("apiserver-server-name-override", "", "It will override the virtual host name of authority")
	cmd.Flags().String("auth-token", "", "Auth token")
	cmd.Flags().String("ca-file", "", "Certificate authority file")

	cmdutil.Must(viper.BindPFlag("apiserver.addr", cmd.Flags().Lookup("apiserver-addr")))
	cmdutil.Must(viper.BindPFlag("apiserver.ca-file", cmd.Flags().Lookup("ca-file")))
	cmdutil.Must(viper.BindPFlag("apiserver.insecure", cmd.Flags().Lookup("apiserver-insecure")))
	cmdutil.Must(viper.BindPFlag("apiserver.insecure-skip-tls-verify", cmd.Flags().Lookup("apiserver-insecure-skip-tls-verify")))
	cmdutil.Must(viper.BindPFlag("apiserver.server-name-override", cmd.Flags().Lookup("apiserver-server-name-override")))
	cmdutil.Must(viper.BindPFlag("auth-token", cmd.Flags().Lookup("auth-token")))
	cmdutil.Must(viper.BindPFlag("ca-file", cmd.Flags().Lookup("ca-file")))

	return cmd
}

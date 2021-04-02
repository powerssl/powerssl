package cmd

import (
	"github.com/spf13/cobra"

	cmdutil "powerssl.dev/common/cmd"

	"powerssl.dev/powerutil/internal"
)

func newCmdVault() *cobra.Command {
	var addr, ca, caKey string

	cmd := &cobra.Command{
		Use:   "vault",
		Short: "Vault migrations",
		Args:  cobra.NoArgs,
		Run: cmdutil.HandleError(func(cmd *cobra.Command, args []string) error {
			return internal.RunVault(addr, ca, caKey)
		}),
	}

	cmd.Flags().StringVar(&addr, "addr", "https://localhost:8200", "Vault address")
	cmd.Flags().StringVar(&ca, "ca", "/etc/powerssl/ca.pem", "Certificate authority file")
	cmd.Flags().StringVar(&caKey, "ca-key", "", "Certificate authority private key file")

	return cmd
}

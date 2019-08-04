package cmd

import (
	"github.com/spf13/cobra"

	"powerssl.dev/powerssl/internal/app/powerutil"
)

func newCmdVault() *cobra.Command {
	var addr, ca, caKey string

	cmd := &cobra.Command{
		Use:   "vault",
		Short: "Vault migrations",
		RunE: func(cmd *cobra.Command, args []string) error {
			return powerutil.RunVault(addr, ca, caKey)
		},
	}

	cmd.Flags().StringVar(&addr, "addr", "https://localhost:8200", "Vault address")
	cmd.Flags().StringVar(&ca, "ca", "/etc/powerssl/ca.pem", "Certificate authority file")
	cmd.Flags().StringVar(&caKey, "ca-key", "", "Certificate authority private key file")

	return cmd
}

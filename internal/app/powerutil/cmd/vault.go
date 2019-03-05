package cmd

import (
	"github.com/spf13/cobra"

	"powerssl.io/internal/app/powerutil"
)

func newCmdVault() *cobra.Command {
	var addr, ca string

	cmd := &cobra.Command{
		Use:   "vault",
		Short: "Vault migrations",
		RunE: func(cmd *cobra.Command, args []string) error {
			return powerutil.RunVault(addr, ca)
		},
	}

	cmd.Flags().StringVar(&addr, "addr", "https://localhost:8200", "Vault address")
	cmd.Flags().StringVar(&ca, "ca", "/etc/powerssl/ca.pem", "Certificate authority file")

	return cmd
}

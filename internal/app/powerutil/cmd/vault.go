package cmd

import (
	"github.com/spf13/cobra"

	"powerssl.io/internal/app/powerutil"
)

func newCmdVault() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vault",
		Short: "Vault migrations",
		RunE: func(cmd *cobra.Command, args []string) error {
			addr := "https://localhost:8200"
			caFile := "/etc/powerssl/ca.pem"

			return powerutil.RunVault(addr, caFile)
		},
	}

	return cmd
}

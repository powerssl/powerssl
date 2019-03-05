package cmd

import (
	"fmt"

	"github.com/cloudflare/cfssl/cli"
	"github.com/spf13/cobra"

	"powerssl.io/internal/pkg/pki"
)

func newCmdCASign() *cobra.Command {
	var ca, caKey, csr string

	cmd := &cobra.Command{
		Use:   "sign",
		Short: "Signs a certificate by a given CA and CA key",
		RunE: func(cmd *cobra.Command, args []string) error {
			csr, err := cli.ReadStdin(csr)
			if err != nil {
				return err
			}

			cert, err := pki.Sign(ca, caKey, string(csr))
			if err != nil {
				return err
			}

			fmt.Println(string(cert))
			return nil
		},
	}

	cmd.Flags().StringVar(&ca, "ca", "", "Certificate authority file")
	cmd.Flags().StringVar(&caKey, "ca-key", "", "Certificate authority key file")
	cmd.Flags().StringVar(&csr, "csr", "", "Certificate signing request")

	return cmd
}

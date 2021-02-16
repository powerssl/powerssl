package cmd

import (
	"github.com/cloudflare/cfssl/cli"
	"github.com/spf13/cobra"

	cmdutil "powerssl.dev/common/cmd"
	"powerssl.dev/powerutil/internal/pki"
)

func newCmdCASign() *cobra.Command {
	var ca, caKey, csr string

	cmd := &cobra.Command{
		Use:   "sign",
		Short: "Signs a certificate by a given CA and CA key",
		Args:  cobra.NoArgs,
		Run: cmdutil.HandleError(func(cmd *cobra.Command, args []string) error {
			var byt, cert []byte
			var err error
			if byt, err = cli.ReadStdin(csr); err != nil {
				return err
			}
			if cert, err = pki.Sign(ca, caKey, string(byt)); err != nil {
				return err
			}
			cmd.Println(string(cert))
			return nil
		}),
	}

	cmd.Flags().StringVar(&ca, "ca", "", "Certificate authority file")
	cmd.Flags().StringVar(&caKey, "ca-key", "", "Certificate authority key file")
	cmd.Flags().StringVar(&csr, "csr", "", "Certificate signing request")

	return cmd
}

package cmd

import (
	"fmt"

	"github.com/cloudflare/cfssl/cli"
	"github.com/cloudflare/cfssl/config"
	"github.com/cloudflare/cfssl/helpers"
	"github.com/cloudflare/cfssl/signer"
	"github.com/cloudflare/cfssl/signer/local"
	"github.com/spf13/cobra"
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

			policy := &config.Signing{
				Default: &config.SigningProfile{
					Expiry: helpers.OneYear,
					CAConstraint: config.CAConstraint{
						IsCA:       true,
						MaxPathLen: 1,
					},
					Usage: []string{
						"digital signature",
						"signing",
						"key encipherment",
						"cert sign",
						"crl sign",
					},
				},
			}

			s, err := local.NewSignerFromFile(ca, caKey, policy)
			if err != nil {
				return err
			}

			req := signer.SignRequest{Request: string(csr)}
			cert, err := s.Sign(req)
			if err != nil {
				return err
			}

			fmt.Print(string(cert))

			return nil
		},
	}

	cmd.Flags().StringVar(&ca, "ca", "", "Certificate authority file")
	cmd.Flags().StringVar(&caKey, "ca-key", "", "Certificate authority key file")
	cmd.Flags().StringVar(&csr, "csr", "", "Certificate signing request")

	return cmd
}

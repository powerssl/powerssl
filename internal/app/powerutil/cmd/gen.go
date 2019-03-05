package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/cloudflare/cfssl/cli/genkey"
	"github.com/cloudflare/cfssl/config"
	"github.com/cloudflare/cfssl/csr"
	"github.com/cloudflare/cfssl/signer"
	"github.com/cloudflare/cfssl/signer/local"
	"github.com/spf13/cobra"
)

func newCmdCAGen() *cobra.Command {
	var ca, caKey, hostname, keyAlgo string
	var keySize int

	cmd := &cobra.Command{
		Use:   "gen",
		Short: "Generate certificate",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := csr.CertificateRequest{
				KeyRequest: &csr.BasicKeyRequest{
					A: keyAlgo,
					S: keySize,
				},
				Hosts: []string{hostname},
			}

			g := &csr.Generator{Validator: genkey.Validator}
			csr, key, err := g.ProcessRequest(&req)
			if err != nil {
				return err
			}

			policy := &config.Signing{Default: config.DefaultConfig()}

			s, err := local.NewSignerFromFile(ca, caKey, policy)
			if err != nil {
				return err
			}

			signReq := signer.SignRequest{Request: string(csr)}
			cert, err := s.Sign(signReq)
			if err != nil {
				return err
			}

			if err := ioutil.WriteFile(fmt.Sprintf("%s.pem", hostname), cert, 0644); err != nil {
				return err
			}
			if err := ioutil.WriteFile(fmt.Sprintf("%s.csr", hostname), csr, 0644); err != nil {
				return err
			}
			if err := ioutil.WriteFile(fmt.Sprintf("%s-key.pem", hostname), key, 0644); err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&ca, "ca", "", "Certificate authority file")
	cmd.Flags().StringVar(&caKey, "ca-key", "", "Certificate authority key file")
	cmd.Flags().StringVar(&hostname, "hostname", "", "Hostname")
	cmd.Flags().StringVar(&keyAlgo, "key-algo", "rsa", "Key algorithm")
	cmd.Flags().IntVar(&keySize, "key-size", 4096, "Key size")

	return cmd
}

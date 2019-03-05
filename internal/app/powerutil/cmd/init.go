package cmd

import (
	"io/ioutil"

	"github.com/cloudflare/cfssl/csr"
	"github.com/cloudflare/cfssl/initca"
	"github.com/spf13/cobra"
)

func newCmdCAInit() *cobra.Command {
	var keyAlgo string
	var keySize int

	cmd := &cobra.Command{
		Use:   "init",
		Short: "Init certificate authority",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := csr.CertificateRequest{
				KeyRequest: &csr.BasicKeyRequest{
					A: keyAlgo,
					S: keySize,
				},
				Names: []csr.Name{
					{
						O: "PowerSSL Root Authority",
					},
				},
			}

			cert, csr, key, err := initca.New(&req)
			if err != nil {
				return err
			}

			if err := ioutil.WriteFile("ca.pem", cert, 0644); err != nil {
				return err
			}
			if err := ioutil.WriteFile("ca.csr", csr, 0644); err != nil {
				return err
			}
			if err := ioutil.WriteFile("ca-key.pem", key, 0644); err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&keyAlgo, "key-algo", "rsa", "Key algorithm")
	cmd.Flags().IntVar(&keySize, "key-size", 4096, "Key size")

	return cmd
}

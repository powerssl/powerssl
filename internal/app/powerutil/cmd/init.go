package cmd

import (
	"io/ioutil"

	"github.com/cloudflare/cfssl/csr"
	"github.com/cloudflare/cfssl/initca"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newCmdCAInit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Init certificate authority",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := csr.CertificateRequest{
				KeyRequest: &csr.BasicKeyRequest{
					A: viper.GetString("key-algo"),
					S: viper.GetInt("key-size"),
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

	cmd.Flags().StringP("key-algo", "", "rsa", "Key algorithm")
	cmd.Flags().IntP("key-size", "", 4096, "Key size")

	viper.BindPFlag("key-algo", cmd.Flags().Lookup("key-algo"))
	viper.BindPFlag("key-size", cmd.Flags().Lookup("key-size"))

	return cmd
}

package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/cloudflare/cfssl/cli"
	"github.com/cloudflare/cfssl/cli/genkey"
	"github.com/cloudflare/cfssl/cli/sign"
	"github.com/cloudflare/cfssl/csr"
	"github.com/cloudflare/cfssl/signer"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newCmdCAGen() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gen",
		Short: "Generate certificate",
		RunE: func(cmd *cobra.Command, args []string) error {
			c := cli.Config{
				CAFile:    viper.GetString("ca"),
				CAKeyFile: viper.GetString("ca-key"),
			}

			req := csr.CertificateRequest{
				KeyRequest: &csr.BasicKeyRequest{
					A: viper.GetString("key-algo"),
					S: viper.GetInt("key-size"),
				},
				Hosts: []string{viper.GetString("hostname")},
			}

			g := &csr.Generator{Validator: genkey.Validator}
			csr, key, err := g.ProcessRequest(&req)
			if err != nil {
				return err
			}

			s, err := sign.SignerFromConfig(c)
			if err != nil {
				return err
			}

			signReq := signer.SignRequest{
				Request: string(csr),
			}

			cert, err := s.Sign(signReq)
			if err != nil {
				return err
			}

			if err := ioutil.WriteFile(fmt.Sprint(viper.GetString("hostname"), ".pem"), cert, 0644); err != nil {
				return err
			}
			if err := ioutil.WriteFile(fmt.Sprint(viper.GetString("hostname"), ".csr"), csr, 0644); err != nil {
				return err
			}
			if err := ioutil.WriteFile(fmt.Sprint(viper.GetString("hostname"), "-key.pem"), key, 0644); err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringP("ca", "", "", "Certificate authority file")
	cmd.Flags().StringP("ca-key", "", "", "Certificate authority key file")
	cmd.Flags().StringP("hostname", "", "", "Hostname")
	cmd.Flags().StringP("key-algo", "", "rsa", "Key algorithm")
	cmd.Flags().IntP("key-size", "", 4096, "Key size")

	viper.BindPFlag("ca", cmd.Flags().Lookup("ca"))
	viper.BindPFlag("ca-key", cmd.Flags().Lookup("ca-key"))
	viper.BindPFlag("hostname", cmd.Flags().Lookup("hostname"))
	viper.BindPFlag("key-algo", cmd.Flags().Lookup("key-algo"))
	viper.BindPFlag("key-size", cmd.Flags().Lookup("key-size"))

	return cmd
}

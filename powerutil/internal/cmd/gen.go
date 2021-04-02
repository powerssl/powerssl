package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"

	cmdutil "powerssl.dev/common/cmd"

	"powerssl.dev/powerutil/internal/pki"
)

func newCmdCAGen() *cobra.Command {
	var ca, caKey, hostname, keyAlgo string
	var keySize int

	cmd := &cobra.Command{
		Use:   "gen",
		Short: "Generate certificate",
		Args:  cobra.NoArgs,
		Run: cmdutil.HandleError(func(cmd *cobra.Command, args []string) error {
			var cert, csr, key []byte
			var err error
			if cert, csr, key, err = pki.Gen(ca, caKey, hostname, keyAlgo, keySize); err != nil {
				return err
			}
			if err = ioutil.WriteFile(fmt.Sprintf("%s.pem", hostname), cert, 0644); err != nil {
				return err
			}
			if err = ioutil.WriteFile(fmt.Sprintf("%s.csr", hostname), csr, 0644); err != nil {
				return err
			}
			if err = ioutil.WriteFile(fmt.Sprintf("%s-key.pem", hostname), key, 0644); err != nil {
				return err
			}
			return nil
		}),
	}

	cmd.Flags().StringVar(&ca, "ca", "", "Certificate authority file")
	cmd.Flags().StringVar(&caKey, "ca-key", "", "Certificate authority key file")
	cmd.Flags().StringVar(&hostname, "hostname", "", "Hostname")
	cmd.Flags().StringVar(&keyAlgo, "key-algo", "rsa", "Key algorithm")
	cmd.Flags().IntVar(&keySize, "key-size", 4096, "Key size")

	return cmd
}

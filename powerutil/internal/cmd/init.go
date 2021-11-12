package cmd

import (
	"io/ioutil"

	"github.com/spangenberg/snakecharmer"
	"github.com/spf13/cobra"

	"powerssl.dev/powerutil/internal/pki"
)

func newCmdCAInit() *cobra.Command {
	var keyAlgo string
	var keySize int

	cmd := &cobra.Command{
		Use:   "init",
		Short: "Init certificate authority",
		Args:  cobra.NoArgs,
		Run: snakecharmer.HandleError(func(cmd *cobra.Command, args []string) error {
			var cert, csr, key []byte
			var err error
			if cert, csr, key, err = pki.Init(keyAlgo, keySize); err != nil {
				return err
			}
			if err = ioutil.WriteFile("ca.pem", cert, 0644); err != nil {
				return err
			}
			if err = ioutil.WriteFile("ca.csr", csr, 0644); err != nil {
				return err
			}
			if err = ioutil.WriteFile("ca-key.pem", key, 0644); err != nil {
				return err
			}
			return nil
		}),
	}

	cmd.Flags().StringVar(&keyAlgo, "key-algo", "rsa", "Key algorithm")
	cmd.Flags().IntVar(&keySize, "key-size", 4096, "Key size")

	return cmd
}

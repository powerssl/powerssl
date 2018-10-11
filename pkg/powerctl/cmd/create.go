package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"powerssl.io/pkg/api"
	"powerssl.io/pkg/powerctl"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a resource.",
	Long: `Create a resource.

Available Commands:
  certificateauthority Create a CertificateAuthority`,
}

var createCertificateCmd = &cobra.Command{
	Use:   "certificate",
	Short: "Create a certificate.",
	Run: func(cmd *cobra.Command, args []string) {
		grpcAddr := viper.GetString("grpcAddr")
		c := powerctl.NewGRPCClient(grpcAddr)

		certificate, err := c.Certificate.Create(context.Background(), &api.Certificate{
			Dnsnames:        []string{"example.com"},
			DigestAlgorithm: "SHA1",
		})
		fmt.Printf("err: %#v\n", err)
		fmt.Printf("certificate: %#v\n", certificate)

		certificates, nextPageToken, _ := c.Certificate.List(context.Background(), 0, "")
		fmt.Printf("nextPageToken: %q", nextPageToken)
		for _, certificate := range certificates {
			fmt.Printf("%#v\n", certificate.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(createCertificateCmd)
}

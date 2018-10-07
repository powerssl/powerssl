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
			TypeMeta: api.TypeMeta{
				APIVersion: "v1",
				Kind:       "Certificate",
			},
			ObjectMeta: api.ObjectMeta{
				Labels: map[string]string{
					"foo": "bar",
					"baz": "boo",
				},
			},
			Spec: api.CertificateSpec{
				CommonName: "rofl",
			},
		})
		fmt.Printf("err: %#v\n", err)
		fmt.Printf("certificate: %#v\n", certificate)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(createCertificateCmd)
}

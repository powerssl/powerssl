package cmd

import (
	"log"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"powerssl.io/pkg/api/v1"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a resource.",
	Long: `Create a resource.

Available Commands:
  certificateauthority Create a CertificateAuthority`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		address := "localhost:8080"
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := api.NewCertificateAuthorityServiceClient(conn)
		log.Println(c)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		certificateAuthority := &api.CertificateAuthority{
			TypeMeta: &api.TypeMeta{
				ApiVersion: "v1",
				Kind:       "CertificateAuthority",
			},
			ObjectMeta: &api.ObjectMeta{
				Labels: map[string]string{
					"foo": "bar",
					"baz": "boo",
				},
			},
			Spec: &api.CertificateAuthoritySpec{
				Vendor: "rofl",
			},
		}
		response, err := c.CreateCertificateAuthority(ctx, &api.CreateCertificateAuthorityRequest{
			CertificateAuthority: certificateAuthority,
		})
		if err != nil {
			log.Fatalf("could not list: %v", err)
		}
		log.Println("RESPONSE:")
		log.Println(response)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}

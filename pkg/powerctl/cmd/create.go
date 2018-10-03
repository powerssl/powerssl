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

		response, err := c.ListCertificateAuthorities(ctx, &api.ListCertificateAuthoritiesRequest{})
		if err != nil {
			log.Fatalf("could not list: %v", err)
		}
		log.Println(response)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}

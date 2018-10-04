package cmd

import (
	"log"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	apiv1 "powerssl.io/pkg/api/v1"
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
		c := apiv1.NewCertificateAuthorityServiceClient(conn)
		log.Println(c)

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		certificateAuthority := &apiv1.CertificateAuthority{
			TypeMeta: &apiv1.TypeMeta{
				ApiVersion: "v1",
				Kind:       "CertificateAuthority",
			},
			ObjectMeta: &apiv1.ObjectMeta{
				Labels: map[string]string{
					"foo": "bar",
					"baz": "boo",
				},
			},
			Spec: &apiv1.CertificateAuthoritySpec{
				Vendor: "rofl",
			},
		}
		n := 10000
		done := make(chan bool)
		for i := 1; i <= n; i++ {
			go func() {
				_, err := c.CreateCertificateAuthority(ctx, &apiv1.CreateCertificateAuthorityRequest{
					CertificateAuthority: certificateAuthority,
				})
				if err != nil {
					log.Fatalf("could not list: %v", err)
				}
				//log.Println("RESPONSE:")
				//log.Println(response)
				done <- true
			}()
		}
		for i := 0; i < n; i++ {
			<-done
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}

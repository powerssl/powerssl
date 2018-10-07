package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-kit/kit/log"
	// "github.com/gogo/protobuf/types"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	// apiv1 "powerssl.io/pkg/api/v1"
	"powerssl.io/pkg/api"

	"powerssl.io/pkg/powerctl"
)

// TODO
const grpcAddr = "localhost:8080"

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a resource.",
	Long: `Create a resource.

Available Commands:
  certificateauthority Create a CertificateAuthority`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v", err)
			os.Exit(1)
		}
		defer conn.Close()
		c := powerctl.NewGRPCClient(conn, log.NewNopLogger())

		v, err := c.CertificateAuthority.Create(context.Background(), &api.CertificateAuthority{})
		fmt.Printf("%+v\n", err)
		fmt.Printf("%+v\n", v)
		fmt.Printf("%#v\n", v)

		// certificateAuthority := &apiv1.CertificateAuthority{
		// 	TypeMeta: &apiv1.TypeMeta{
		// 		ApiVersion: "v1",
		// 		Kind:       "CertificateAuthority",
		// 	},
		// 	ObjectMeta: &apiv1.ObjectMeta{
		// 		Labels: map[string]string{
		// 			"foo": "bar",
		// 			"baz": "boo",
		// 		},
		// 		CreationTimestamp: types.TimestampNow(),
		// 	},
		// 	Spec: &apiv1.CertificateAuthoritySpec{
		// 		Vendor: "rofl",
		// 	},
		// }
		// n := 1
		// done := make(chan bool)
		// for i := 1; i <= n; i++ {
		// 	go func() {
		// 		response, err := c.CreateCertificateAuthority(ctx, &apiv1.CreateCertificateAuthorityRequest{
		// 			CertificateAuthority: certificateAuthority,
		// 		})
		// 		if err != nil {
		// 			fmt.Fatalf("could not list: %v", err)
		// 		}
		// 		fmt.Println("RESPONSE:")
		// 		fmt.Println(response)
		// 		done <- true
		// 	}()
		// }
		// for i := 0; i < n; i++ {
		// 	<-done
		// }
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}

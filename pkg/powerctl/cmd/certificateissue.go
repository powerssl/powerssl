package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"powerssl.io/pkg/api"
)

var (
	certificateIssueCmd = &cobra.Command{
		Use:   "certificateissue",
		Short: "Certificate issue resource",
		Long:  `Certificate issue resource.`,
	}

	createCertificateIssueCmd = &cobra.Command{
		Use:   "create",
		Short: "Create Certificate issue.",
		Args:  validateParentArg("certificates"),
		Run:   createCertificateIssue,
	}

	deleteCertificateIssueCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete Certificate issue.",
		Args:  validateNameArg,
		Run:   deleteCertificateIssue,
	}

	getCertificateIssueCmd = &cobra.Command{
		Use:   "get",
		Short: "Get Certificate issue.",
		Args:  validateNameArg,
		Run:   getCertificateIssue,
	}

	listCertificateIssueCmd = &cobra.Command{
		Use:   "list",
		Short: "List Certificates.",
		Args:  validateParentArg("certificates"),
		Run:   listCertificateIssue,
	}

	updateCertificateIssueCmd = &cobra.Command{
		Use:   "update",
		Short: "Update Certificate issue.",
		Args:  validateNameArg,
		Run:   updateCertificateIssue,
	}
)

func init() {
	rootCmd.AddCommand(certificateIssueCmd)
	certificateIssueCmd.AddCommand(
		createCertificateIssueCmd,
		deleteCertificateIssueCmd,
		getCertificateIssueCmd,
		listCertificateIssueCmd,
		updateCertificateIssueCmd)
}

func createCertificateIssue(cmd *cobra.Command, args []string) {
	parent := args[0]
	certificateIssue := makeCertificateIssue()
	client := newGRPCClient()
	createResource(func() (interface{}, error) {
		return client.CertificateIssue.Create(context.Background(), parent, certificateIssue)
	})
}

func deleteCertificateIssue(cmd *cobra.Command, args []string) {
	name := args[0]
	client := newGRPCClient()
	deleteResource(func() error {
		return client.CertificateIssue.Delete(context.Background(), name)
	})
}

func getCertificateIssue(cmd *cobra.Command, args []string) {
	name := args[0]
	client := newGRPCClient()
	getResource(func() (interface{}, error) {
		return client.CertificateIssue.Get(context.Background(), name)
	})
}

func listCertificateIssue(cmd *cobra.Command, args []string) {
	parent := args[0]
	client := newGRPCClient()
	listResource(func(pageToken string) (interface{}, string, error) {
		return client.CertificateIssue.List(context.Background(), parent, 0, pageToken)
	})
}

func updateCertificateIssue(cmd *cobra.Command, args []string) {
	name := args[0]
	certificateIssue := makeCertificateIssue()
	client := newGRPCClient()
	updateResource(func() (interface{}, error) {
		return client.CertificateIssue.Update(context.Background(), name, certificateIssue)
	})
}

func makeCertificateIssue() *api.CertificateIssue {
	return &api.CertificateIssue{}
}

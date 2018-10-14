package cmd

import (
	"context"
	"strings"

	"github.com/spf13/cobra"

	"powerssl.io/pkg/api"
)

var (
	AutoRenew       bool
	DNSNames        string
	DigestAlgorithm string
	KeyAlgorithm    string
	KeySize         int

	certificateCmd = &cobra.Command{
		Use:   "certificate",
		Short: "Certificate resource",
		Long:  `Certificate resource.`,
	}

	createCertificateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create Certificate.",
		Args:  cobra.NoArgs,
		Run:   createCertificate,
	}

	deleteCertificateCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete Certificate.",
		Args:  validateNameArg,
		Run:   deleteCertificate,
	}

	getCertificateCmd = &cobra.Command{
		Use:   "get",
		Short: "Get Certificate.",
		Args:  validateNameArg,
		Run:   getCertificate,
	}

	listCertificateCmd = &cobra.Command{
		Use:   "list",
		Short: "List Certificates.",
		Args:  cobra.NoArgs,
		Run:   listCertificate,
	}

	updateCertificateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update Certificate.",
		Args:  validateNameArg,
		Run:   updateCertificate,
	}
)

func init() {
	createCertificateCmd.Flags().StringVarP(&DNSNames, "dns-names", "", "", "DNS name for the certificate (seperated by \",\")")
	createCertificateCmd.Flags().StringVarP(&KeyAlgorithm, "key-algorithm", "", "", "Key algorithm ...")
	createCertificateCmd.Flags().IntVarP(&KeySize, "key-size", "", 0, "Key size ...")
	createCertificateCmd.Flags().StringVarP(&DigestAlgorithm, "digest-algorithm", "", "", "Digest algorithm ...")
	createCertificateCmd.Flags().BoolVarP(&AutoRenew, "auto-renew", "", false, "Auto renew ...")
	createCertificateCmd.MarkFlagRequired("dns-names")
	createCertificateCmd.MarkFlagRequired("key-algorithm")
	createCertificateCmd.MarkFlagRequired("key-size")
	createCertificateCmd.MarkFlagRequired("digest-algorithm")

	updateCertificateCmd.Flags().StringVarP(&DNSNames, "dns-names", "", "", "DNS name for the certificate (seperated by \",\")")
	updateCertificateCmd.Flags().StringVarP(&KeyAlgorithm, "key-algorithm", "", "", "Key algorithm ...")
	updateCertificateCmd.Flags().IntVarP(&KeySize, "key-size", "", 0, "Key size ...")
	updateCertificateCmd.Flags().StringVarP(&DigestAlgorithm, "digest-algorithm", "", "", "Digest algorithm ...")
	updateCertificateCmd.Flags().BoolVarP(&AutoRenew, "auto-renew", "", false, "Auto renew ...")

	rootCmd.AddCommand(certificateCmd)
	certificateCmd.AddCommand(
		createCertificateCmd,
		deleteCertificateCmd,
		getCertificateCmd,
		listCertificateCmd,
		updateCertificateCmd)
}

func createCertificate(cmd *cobra.Command, args []string) {
	client := newGRPCClient()
	createResource(func() (interface{}, error) {
		return client.Certificate.Create(context.Background(), makeCertificate())
	})
}

func deleteCertificate(cmd *cobra.Command, args []string) {
	name := args[0]
	client := newGRPCClient()
	deleteResource(func() error {
		return client.Certificate.Delete(context.Background(), name)
	})
}

func getCertificate(cmd *cobra.Command, args []string) {
	name := args[0]
	client := newGRPCClient()
	getResource(func() (interface{}, error) {
		return client.Certificate.Get(context.Background(), name)
	})
}

func listCertificate(cmd *cobra.Command, args []string) {
	client := newGRPCClient()
	listResource(func(pageToken string) (interface{}, string, error) {
		return client.Certificate.List(context.Background(), 0, pageToken)
	})
}

func updateCertificate(cmd *cobra.Command, args []string) {
	name := args[0]
	client := newGRPCClient()
	updateResource(func() (interface{}, error) {
		return client.Certificate.Update(context.Background(), name, makeCertificate())
	})
}

func makeCertificate() *api.Certificate {
	return &api.Certificate{
		Dnsnames:        strings.Split(DNSNames, ","),
		KeyAlgorithm:    KeyAlgorithm,
		KeySize:         int32(KeySize),
		DigestAlgorithm: DigestAlgorithm,
		AutoRenew:       AutoRenew,
	}
}

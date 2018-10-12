package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"

	"powerssl.io/pkg/api"
	"powerssl.io/pkg/powerctl"
)

func newGRPCClient() *powerctl.GrpcClient {
	grpcAddr := viper.GetString("grpcAddr")
	return powerctl.NewGRPCClient(grpcAddr)
}

// certificateCmd represents the certificate command
var certificateCmd = &cobra.Command{
	Use:   "certificate",
	Short: "Certificate resource",
	Long:  `Certificate resource.`,
}

func printResource(resource interface{}) {
	byt, err := yaml.Marshal(resource)
	if err != nil {
		er(err)
	}
	fmt.Println(string(byt))
}

var (
	AutoRenew       bool
	DNSNames        string
	DigestAlgorithm string
	KeyAlgorithm    string
	KeySize         int
	Name            string
	PageSize        int
)

var createCertificateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a certificate.",
	Run: func(cmd *cobra.Command, args []string) {
		client := newGRPCClient()
		certificate, err := client.Certificate.Create(context.Background(), &api.Certificate{
			Dnsnames:        strings.Split(DNSNames, ","),
			KeyAlgorithm:    KeyAlgorithm,
			KeySize:         int32(KeySize),
			DigestAlgorithm: DigestAlgorithm,
			AutoRenew:       AutoRenew,
		})
		if err != nil {
			er(err)
		}
		printResource(certificate)
	},
}

var deleteCertificateCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete Certificate.",
	Run: func(cmd *cobra.Command, args []string) {
		client := newGRPCClient()
		if err := client.Certificate.Delete(context.Background(), Name); err != nil {
			er(err)
		}
	},
}

var getCertificateCmd = &cobra.Command{
	Use:   "get",
	Short: "Get Certificate.",
	Run: func(cmd *cobra.Command, args []string) {
		client := newGRPCClient()
		certificate, err := client.Certificate.Get(context.Background(), Name)
		if err != nil {
			er(err)
		}
		printResource(certificate)
	},
}

var listCertificateCmd = &cobra.Command{
	Use:   "list",
	Short: "List Certificates.",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			certificates []*api.Certificate
			pageToken    string
		)
		client := newGRPCClient()
		for {
			certs, nextPageToken, err := client.Certificate.List(context.Background(), PageSize, pageToken)
			if err != nil {
				er(err)
			}
			for _, cert := range certs {
				certificates = append(certificates, cert)
			}
			if nextPageToken == "" {
				break
			} else {
				pageToken = nextPageToken
			}
		}
		printResource(certificates)
	},
}

var updateCertificateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Certificate.",
	Run: func(cmd *cobra.Command, args []string) {
		client := newGRPCClient()
		certificate, err := client.Certificate.Update(context.Background(), Name, &api.Certificate{
			Dnsnames:        strings.Split(DNSNames, ","),
			KeyAlgorithm:    KeyAlgorithm,
			KeySize:         int32(KeySize),
			DigestAlgorithm: DigestAlgorithm,
			AutoRenew:       AutoRenew,
		})
		if err != nil {
			er(err)
		}
		printResource(certificate)
	},
}

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

	deleteCertificateCmd.Flags().StringVarP(&Name, "name", "", "", "Name ...")
	deleteCertificateCmd.MarkFlagRequired("name")

	getCertificateCmd.Flags().StringVarP(&Name, "name", "", "", "Name ...")
	getCertificateCmd.MarkFlagRequired("name")

	listCertificateCmd.Flags().IntVarP(&PageSize, "page-size", "", 0, "Page size ...")

	updateCertificateCmd.Flags().StringVarP(&Name, "name", "", "", "Name ...")
	updateCertificateCmd.Flags().StringVarP(&DNSNames, "dns-names", "", "", "DNS name for the certificate (seperated by \",\")")
	updateCertificateCmd.Flags().StringVarP(&KeyAlgorithm, "key-algorithm", "", "", "Key algorithm ...")
	updateCertificateCmd.Flags().IntVarP(&KeySize, "key-size", "", 0, "Key size ...")
	updateCertificateCmd.Flags().StringVarP(&DigestAlgorithm, "digest-algorithm", "", "", "Digest algorithm ...")
	updateCertificateCmd.Flags().BoolVarP(&AutoRenew, "auto-renew", "", false, "Auto renew ...")
	updateCertificateCmd.MarkFlagRequired("name")

	rootCmd.AddCommand(certificateCmd)
	certificateCmd.AddCommand(
		createCertificateCmd,
		deleteCertificateCmd,
		getCertificateCmd,
		listCertificateCmd,
		updateCertificateCmd)
}

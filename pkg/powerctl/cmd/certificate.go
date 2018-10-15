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

	createCertificateCmd = &cobra.Command{
		Use:   "certificate",
		Short: "Create Certificate",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			certificate := &api.Certificate{}
			if Filename != "" {
				loadResource(Filename, certificate)
			} else {
				certificate = makeCertificate()
			}
			createCertificate(certificate)
		},
	}

	deleteCertificateCmd = &cobra.Command{
		Use:   "certificate",
		Short: "Delete Certificate",
		Args:  validateNameArg,
		Run: func(cmd *cobra.Command, args []string) {
			deleteCertificate(nameArg("certificates", args[0]))
		},
	}

	getCertificateCmd = &cobra.Command{
		Use:     "certificate",
		Aliases: []string{"certificates"},
		Short:   "Get Certificate",
		Example: `  powerctl get certificate       List all certificates
  powerctl get certificate 42    Get an certificate
  powerctl get certificates/42   Get an certificate`,
		Args: cobra.RangeArgs(0, 1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 1 {
				getCertificate(nameArg("certificates", args[0]))
			} else {
				listCertificate()
			}
		},
	}

	updateCertificateCmd = &cobra.Command{
		Use:   "certificate",
		Short: "Update Certificate",
		Args:  validateNameArg,
		Run: func(cmd *cobra.Command, args []string) {
			certificate := &api.Certificate{}
			if Filename != "" {
				loadResource(Filename, certificate)
			} else {
				certificate = makeCertificate()
			}
			updateCertificate(nameArg("certificates", args[0]), certificate)
		},
	}
)

func init() {
	createCertificateCmd.Flags().StringVarP(&Filename, "filename", "f", "", "Filename to file to use to create the certificate")
	createCertificateCmd.Flags().StringVarP(&DNSNames, "dns-names", "", "", "DNS name for the certificate (seperated by \",\")")
	createCertificateCmd.Flags().StringVarP(&KeyAlgorithm, "key-algorithm", "", "", "Key algorithm ...")
	createCertificateCmd.Flags().IntVarP(&KeySize, "key-size", "", 0, "Key size ...")
	createCertificateCmd.Flags().StringVarP(&DigestAlgorithm, "digest-algorithm", "", "", "Digest algorithm ...")
	createCertificateCmd.Flags().BoolVarP(&AutoRenew, "auto-renew", "", false, "Auto renew ...")

	updateCertificateCmd.Flags().StringVarP(&Filename, "filename", "f", "", "Filename to file to use to update the certificate")
	updateCertificateCmd.Flags().StringVarP(&DNSNames, "dns-names", "", "", "DNS name for the certificate (seperated by \",\")")
	updateCertificateCmd.Flags().StringVarP(&KeyAlgorithm, "key-algorithm", "", "", "Key algorithm ...")
	updateCertificateCmd.Flags().IntVarP(&KeySize, "key-size", "", 0, "Key size ...")
	updateCertificateCmd.Flags().StringVarP(&DigestAlgorithm, "digest-algorithm", "", "", "Digest algorithm ...")
	updateCertificateCmd.Flags().BoolVarP(&AutoRenew, "auto-renew", "", false, "Auto renew ...")

	createCmd.AddCommand(createCertificateCmd)
	deleteCmd.AddCommand(deleteCertificateCmd)
	getCmd.AddCommand(getCertificateCmd)
	updateCmd.AddCommand(updateCertificateCmd)
}

func createCertificate(certificate *api.Certificate) {
	client := newGRPCClient()
	createResource(func() (interface{}, error) {
		return client.Certificate.Create(context.Background(), certificate)
	})
}

func deleteCertificate(name string) {
	client := newGRPCClient()
	deleteResource(func() error {
		return client.Certificate.Delete(context.Background(), name)
	})
}

func getCertificate(name string) {
	client := newGRPCClient()
	getResource(func() (interface{}, error) {
		return client.Certificate.Get(context.Background(), name)
	})
}

func listCertificate() {
	client := newGRPCClient()
	listResource(func(pageToken string) (interface{}, string, error) {
		return client.Certificate.List(context.Background(), 0, pageToken)
	})
}

func updateCertificate(name string, certificate *api.Certificate) {
	client := newGRPCClient()
	updateResource(func() (interface{}, error) {
		return client.Certificate.Update(context.Background(), name, certificate)
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

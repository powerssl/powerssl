package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"powerssl.io/pkg/apiserver/api"
	apiserverclient "powerssl.io/pkg/apiserver/client"
)

var Certificate certificate

type CertificateSpec struct {
	Dnsnames        []string `json:"dnsnames,omitempty"        yaml:"dnsnames,omitempty"`
	KeyAlgorithm    string   `json:"keyAlgorithm,omitempty"    yaml:"keyAlgorithm,omitempty"`
	KeySize         int32    `json:"keySize,omitempty"         yaml:"keySize,omitempty"`
	DigestAlgorithm string   `json:"digestAlgorithm,omitempty" yaml:"digestAlgorithm,omitempty"`
	AutoRenew       bool     `json:"autoRenew,omitempty"       yaml:"autoRenew,omitempty"`
}

type certificate struct{}

func (r certificate) Create(client *apiserverclient.GRPCClient, resource *Resource) (*Resource, error) {
	spec := resource.Spec.(*CertificateSpec)
	certificate := &api.Certificate{
		Dnsnames:        spec.Dnsnames,
		KeyAlgorithm:    spec.KeyAlgorithm,
		KeySize:         spec.KeySize,
		DigestAlgorithm: spec.DigestAlgorithm,
		AutoRenew:       spec.AutoRenew,
	}
	certificate, err := client.Certificate.Create(context.Background(), certificate)
	if err != nil {
		return nil, err
	}
	return r.Encode(certificate), nil
}

func (r certificate) Delete(client *apiserverclient.GRPCClient, name string) error {
	return client.Certificate.Delete(context.Background(), fmt.Sprintf("certificates/%s", name))
}

func (r certificate) Encode(certificate *api.Certificate) *Resource {
	uid := strings.Split(certificate.Name, "/")[1]
	return &Resource{
		Kind: "certificate",
		Meta: &ResourceMeta{
			UID:        uid,
			CreateTime: certificate.CreateTime,
			UpdateTime: certificate.UpdateTime,
		},
		Spec: &CertificateSpec{
			Dnsnames:        certificate.Dnsnames,
			KeyAlgorithm:    certificate.KeyAlgorithm,
			KeySize:         certificate.KeySize,
			DigestAlgorithm: certificate.DigestAlgorithm,
			AutoRenew:       certificate.AutoRenew,
		},
	}
}

func (r certificate) Get(client *apiserverclient.GRPCClient, name string) (*Resource, error) {
	certificate, err := client.Certificate.Get(context.Background(), fmt.Sprintf("certificates/%s", name))
	if err != nil {
		return nil, err
	}
	return r.Encode(certificate), nil
}

func (r certificate) List(client *apiserverclient.GRPCClient) ([]*Resource, error) {
	return listResource(func(pageToken string) ([]*Resource, string, error) {
		certificates, nextPageToken, err := client.Certificate.List(context.Background(), 0, pageToken)
		if err != nil {
			return nil, nextPageToken, err
		}
		a := make([]*Resource, len(certificates))
		for i, certificate := range certificates {
			a[i] = r.Encode(certificate)
		}
		return a, nextPageToken, nil
	})
}

func (r certificate) Spec() interface{} {
	return new(CertificateSpec)
}

var (
	AutoRenew       bool
	DNSNames        string
	DigestAlgorithm string
	KeyAlgorithm    string
	KeySize         int
)

var createCertificateCmd = &cobra.Command{
	Use:   "certificate",
	Short: "Create Certificate",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := NewGRPCClient()
		if err != nil {
			er(err)
		}
		certificate := &api.Certificate{
			Dnsnames:        strings.Split(DNSNames, ","),
			KeyAlgorithm:    KeyAlgorithm,
			KeySize:         int32(KeySize),
			DigestAlgorithm: DigestAlgorithm,
			AutoRenew:       AutoRenew,
		}
		certificate, err = client.Certificate.Create(context.Background(), certificate)
		if err != nil {
			er(err)
		}
		pr(Certificate.Encode(certificate))
	},
}

func init() {
	Resources.Add(Certificate, "cert", "c")

	createCertificateCmd.Flags().StringVarP(&Filename, "filename", "f", "", "Filename to file to use to create the certificate")
	createCertificateCmd.Flags().StringVarP(&DNSNames, "dns-names", "", "", "DNS name for the certificate (seperated by \",\")")
	createCertificateCmd.Flags().StringVarP(&KeyAlgorithm, "key-algorithm", "", "", "Key algorithm ...")
	createCertificateCmd.Flags().IntVarP(&KeySize, "key-size", "", 0, "Key size ...")
	createCertificateCmd.Flags().StringVarP(&DigestAlgorithm, "digest-algorithm", "", "", "Digest algorithm ...")
	createCertificateCmd.Flags().BoolVarP(&AutoRenew, "auto-renew", "", false, "Auto renew ...")

	createCmd.AddCommand(createCertificateCmd)
}

package resource

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"

	apiv1 "powerssl.dev/api/apiserver/v1"
	cmdutil "powerssl.dev/common/cmd"
	"powerssl.dev/sdk/apiserver"

	"powerssl.dev/powerctl/internal"
)

type certificate struct{}

func (r certificate) Create(client *apiserver.Client, resource *Resource) (*Resource, error) {
	spec := resource.Spec.(*certificateSpec)
	certificate := &apiv1.Certificate{
		Dnsnames:        spec.Dnsnames,
		KeyAlgorithm:    apiv1.KeyAlgorithm(apiv1.KeyAlgorithm_value[spec.KeyAlgorithm]),
		KeySize:         spec.KeySize,
		DigestAlgorithm: apiv1.DigestAlgorithm(apiv1.DigestAlgorithm_value[spec.DigestAlgorithm]),
		AutoRenew:       spec.AutoRenew,
	}
	certificate, err := client.Certificate.Create(context.Background(), &apiv1.CreateCertificateRequest{
		Certificate: certificate,
	})
	if err != nil {
		return nil, err
	}
	return r.Encode(certificate), nil
}

func (r certificate) Delete(client *apiserver.Client, name string) error {
	_, err := client.Certificate.Delete(context.Background(), &apiv1.DeleteCertificateRequest{
		Name: fmt.Sprintf("certificates/%s", name),
	})
	return err
}

func (r certificate) Encode(certificate *apiv1.Certificate) *Resource {
	uid := strings.Split(certificate.GetName(), "/")[1]
	return &Resource{
		Kind: "certificate",
		Meta: &resourceMeta{
			UID:        uid,
			CreateTime: certificate.GetCreateTime().AsTime(),
			UpdateTime: certificate.GetUpdateTime().AsTime(),
		},
		Spec: &certificateSpec{
			Dnsnames:        certificate.GetDnsnames(),
			KeyAlgorithm:    certificate.GetKeyAlgorithm().String(),
			KeySize:         certificate.GetKeySize(),
			DigestAlgorithm: certificate.GetDigestAlgorithm().String(),
			AutoRenew:       certificate.GetAutoRenew(),
		},
	}
}

func (r certificate) Get(client *apiserver.Client, name string) (*Resource, error) {
	certificate, err := client.Certificate.Get(context.Background(), &apiv1.GetCertificateRequest{
		Name: fmt.Sprintf("certificates/%s", name),
	})
	if err != nil {
		return nil, err
	}
	return r.Encode(certificate), nil
}

func (r certificate) List(client *apiserver.Client) ([]*Resource, error) {
	return listResource(func(pageToken string) ([]*Resource, string, error) {
		response, err := client.Certificate.List(context.Background(), &apiv1.ListCertificatesRequest{
			PageToken: pageToken,
			PageSize:  0,
		})
		if err != nil {
			return nil, "", err
		}
		a := make([]*Resource, len(response.GetCertificates()))
		for i, certificate := range response.GetCertificates() {
			a[i] = r.Encode(certificate)
		}
		return a, response.GetNextPageToken(), nil
	})
}

func (r certificate) Spec() interface{} {
	return new(certificateSpec)
}

func (r certificate) Columns(resource *Resource) ([]string, []string) {
	spec := resource.Spec.(*certificateSpec)
	return []string{
			"DNS NAMES",
			"KEY ALGORITHM",
			"KEY SIZE",
			"DIGEST ALGORITHM",
			"AUTO RENEW",
		}, []string{
			strings.Join(spec.Dnsnames, ", "),
			fmt.Sprint(spec.KeyAlgorithm),
			fmt.Sprint(spec.KeySize),
			fmt.Sprint(spec.DigestAlgorithm),
			fmt.Sprint(spec.AutoRenew),
		}
}

func (r certificate) Describe(_ *apiserver.Client, resource *Resource, output io.Writer) (err error) {
	spec := resource.Spec.(*certificateSpec)
	w := tabwriter.NewWriter(output, 0, 0, 1, ' ', tabwriter.TabIndent)
	_, _ = fmt.Fprintln(w, fmt.Sprintf("UID:\t%s", resource.Meta.UID))
	_, _ = fmt.Fprintln(w, fmt.Sprintf("Create Time:\t%s", resource.Meta.CreateTime))
	_, _ = fmt.Fprintln(w, fmt.Sprintf("Update Time:\t%s", resource.Meta.UpdateTime))
	_, _ = fmt.Fprintln(w, fmt.Sprintf("DNS Names:\t%s", strings.Join(spec.Dnsnames, ",")))
	_, _ = fmt.Fprintln(w, fmt.Sprintf("Key Algorithm:\t%s", fmt.Sprint(spec.KeyAlgorithm)))
	_, _ = fmt.Fprintln(w, fmt.Sprintf("Key Size:\t%s", fmt.Sprint(spec.KeySize)))
	_, _ = fmt.Fprintln(w, fmt.Sprintf("Digest Algorithm:\t%s", fmt.Sprint(spec.DigestAlgorithm)))
	_, _ = fmt.Fprintln(w, fmt.Sprintf("Auto Renew:\t%v", spec.AutoRenew))
	return w.Flush()
}

type certificateSpec struct {
	Dnsnames        []string `json:"dnsnames,omitempty"        yaml:"dnsnames,omitempty"`
	KeyAlgorithm    string   `json:"keyAlgorithm,omitempty"    yaml:"keyAlgorithm,omitempty"`
	KeySize         int32    `json:"keySize,omitempty"         yaml:"keySize,omitempty"`
	DigestAlgorithm string   `json:"digestAlgorithm,omitempty" yaml:"digestAlgorithm,omitempty"`
	AutoRenew       bool     `json:"autoRenew,omitempty"       yaml:"autoRenew,omitempty"`
}

func NewCmdCreateCertificate() *cobra.Command {
	var client *apiserver.Client
	var (
		autoRenew       bool
		dnsNames        string
		digestAlgorithm string
		keyAlgorithm    string
		keySize         int
	)

	cmd := &cobra.Command{
		Use:   "certificate",
		Short: "Create Certificate",
		Args:  cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) (err error) {
			client, err = internal.NewGRPCClient()
			return err
		},
		Run: cmdutil.HandleError(func(cmd *cobra.Command, args []string) (err error) {
			apiCertificate := &apiv1.Certificate{
				Dnsnames:        strings.Split(dnsNames, ","),
				KeyAlgorithm:    apiv1.KeyAlgorithm(apiv1.KeyAlgorithm_value[keyAlgorithm]),
				KeySize:         int32(keySize),
				DigestAlgorithm: apiv1.DigestAlgorithm(apiv1.DigestAlgorithm_value[digestAlgorithm]),
				AutoRenew:       autoRenew,
			}
			if apiCertificate, err = client.Certificate.Create(context.Background(), &apiv1.CreateCertificateRequest{
				Certificate: apiCertificate,
			}); err != nil {
				return err
			}
			return FormatResource(certificate{}.Encode(apiCertificate), os.Stdout)
		}),
	}

	cmd.Flags().BoolVar(&autoRenew, "auto-renew", false, "Auto renew ...")
	cmd.Flags().IntVar(&keySize, "key-size", 0, "Key size ...")
	cmd.Flags().StringVar(&digestAlgorithm, "digest-algorithm", "", "Digest algorithm ...")
	cmd.Flags().StringVar(&dnsNames, "dns-names", "", "DNS name for the certificate (seperated by \",\")")
	cmd.Flags().StringVar(&keyAlgorithm, "key-algorithm", "", "Key algorithm ...")

	return cmd
}

func init() {
	resources.Add(certificate{}, "cert", "c")
}

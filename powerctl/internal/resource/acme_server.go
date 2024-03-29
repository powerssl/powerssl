package resource

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"

	cmdutil "powerssl.dev/common/cmd"
	"powerssl.dev/sdk/apiserver"
	"powerssl.dev/sdk/apiserver/api"

	"powerssl.dev/powerctl/internal"
)

type acmeServer struct{}

func (r acmeServer) Create(client *apiserver.Client, resource *Resource) (*Resource, error) {
	spec := resource.Spec.(*acmeServerSpec)
	acmeServer := &api.ACMEServer{
		DirectoryURL:    spec.DirectoryURL,
		DisplayName:     spec.DisplayName,
		IntegrationName: spec.IntegrationName,
	}
	acmeServer, err := client.ACMEServer.Create(context.Background(), acmeServer)
	if err != nil {
		return nil, err
	}
	return r.Encode(acmeServer), nil
}

func (r acmeServer) Delete(client *apiserver.Client, name string) error {
	return client.ACMEServer.Delete(context.Background(), fmt.Sprintf("acmeServers/%s", name))
}

func (r acmeServer) Encode(acmeServer *api.ACMEServer) *Resource {
	uid := strings.Split(acmeServer.Name, "/")[1]
	return &Resource{
		Kind: "acmeserver",
		Meta: &resourceMeta{
			UID:        uid,
			CreateTime: acmeServer.CreateTime,
			UpdateTime: acmeServer.UpdateTime,
		},
		Spec: &acmeServerSpec{
			DisplayName:     acmeServer.DisplayName,
			DirectoryURL:    acmeServer.DirectoryURL,
			IntegrationName: acmeServer.IntegrationName,
		},
	}
}

func (r acmeServer) Get(client *apiserver.Client, name string) (*Resource, error) {
	acmeServer, err := client.ACMEServer.Get(context.Background(), fmt.Sprintf("acmeServers/%s", name))
	if err != nil {
		return nil, err
	}
	return r.Encode(acmeServer), nil
}

func (r acmeServer) List(client *apiserver.Client) ([]*Resource, error) {
	return listResource(func(pageToken string) ([]*Resource, string, error) {
		acmeServers, nextPageToken, err := client.ACMEServer.List(context.Background(), 0, pageToken)
		if err != nil {
			return nil, nextPageToken, err
		}
		a := make([]*Resource, len(acmeServers))
		for i, acmeServer := range acmeServers {
			a[i] = r.Encode(acmeServer)
		}
		return a, nextPageToken, nil
	})
}

func (r acmeServer) Spec() interface{} {
	return new(acmeServerSpec)
}

func (r acmeServer) Columns(resource *Resource) ([]string, []string) {
	spec := resource.Spec.(*acmeServerSpec)
	return []string{
			"DISPLAY NAME",
			"DIRECTORY URL",
			"INTEGRATION NAME",
		}, []string{
			fmt.Sprint(spec.DisplayName),
			fmt.Sprint(spec.DirectoryURL),
			fmt.Sprint(spec.IntegrationName),
		}
}

func (r acmeServer) Describe(_ *apiserver.Client, resource *Resource, output io.Writer) (err error) {
	spec := resource.Spec.(*acmeServerSpec)
	w := tabwriter.NewWriter(output, 0, 0, 1, ' ', tabwriter.TabIndent)
	_, _ = fmt.Fprintln(w, fmt.Sprintf("UID:\t%s", resource.Meta.UID))
	_, _ = fmt.Fprintln(w, fmt.Sprintf("Create Time:\t%s", resource.Meta.CreateTime))
	_, _ = fmt.Fprintln(w, fmt.Sprintf("Update Time:\t%s", resource.Meta.UpdateTime))
	_, _ = fmt.Fprintln(w, fmt.Sprintf("Display Name:\t%s", spec.DisplayName))
	_, _ = fmt.Fprintln(w, fmt.Sprintf("Directory URL:\t%s", spec.DirectoryURL))
	_, _ = fmt.Fprintln(w, fmt.Sprintf("Integration Name:\t%s", spec.IntegrationName))
	return w.Flush()
}

type acmeServerSpec struct {
	DisplayName     string `json:"displayName,omitempty"     yaml:"displayName,omitempty"`
	DirectoryURL    string `json:"directoryURL,omitempty"    yaml:"directoryURL,omitempty"`
	IntegrationName string `json:"integrationName,omitempty" yaml:"integrationName,omitempty"`
}

func NewCmdCreateACMEServer() *cobra.Command {
	var client *apiserver.Client
	var (
		directoryURL       string
		displayName        string
		integrationName    string
		letsEncrypt        bool
		letsEncryptStaging bool
	)

	cmd := &cobra.Command{
		Use:     "acmeserver",
		Aliases: []string{"acmeServer"},
		Short:   "Create ACME server",
		Args:    cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) (err error) {
			client, err = internal.NewGRPCClient()
			return err
		},
		Run: cmdutil.HandleError(func(cmd *cobra.Command, args []string) (err error) {
			var apiACMEServer *api.ACMEServer
			if letsEncrypt {
				apiACMEServer = &api.ACMEServer{
					DirectoryURL:    "https://acme-v02.api.letsencrypt.org/directory",
					DisplayName:     "Let's Encrypt",
					IntegrationName: "acme",
				}
			} else if letsEncryptStaging {
				apiACMEServer = &api.ACMEServer{
					DirectoryURL:    "https://acme-staging-v02.api.letsencrypt.org/directory",
					DisplayName:     "Let's Encrypt Staging",
					IntegrationName: "acme",
				}
			} else {
				apiACMEServer = &api.ACMEServer{
					DirectoryURL:    directoryURL,
					DisplayName:     displayName,
					IntegrationName: integrationName,
				}
			}
			if apiACMEServer, err = client.ACMEServer.Create(context.Background(), apiACMEServer); err != nil {
				return err
			}
			return FormatResource(acmeServer{}.Encode(apiACMEServer), os.Stdout)
		}),
	}

	cmd.Flags().BoolVar(&letsEncrypt, "letsencrypt", false, "Let's Encrypt defaults")
	cmd.Flags().BoolVar(&letsEncryptStaging, "letsencrypt-staging", false, "Let's Encrypt staging defaults")
	cmd.Flags().StringVar(&directoryURL, "directory-url", "", "Directory URL")
	cmd.Flags().StringVar(&displayName, "display-name", "", "Display name")
	cmd.Flags().StringVar(&integrationName, "integration-name", "", "Integration name")

	return cmd
}

func init() {
	resources.Add(acmeServer{}, "as")
}

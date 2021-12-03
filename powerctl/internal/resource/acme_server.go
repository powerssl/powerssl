package resource

import (
	"context"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"

	apiv1 "powerssl.dev/api/apiserver/v1"
	"powerssl.dev/sdk/apiserver"

	"powerssl.dev/powerctl/internal"
)

func NewCmdCreateACMEServer() *cobra.Command {
	var (
		directoryURL       string
		displayName        string
		integrationName    string
		letsEncrypt        bool
		letsEncryptStaging bool
	)

	cmd := internal.CmdWithClient(&cobra.Command{
		Use:     "acmeserver",
		Aliases: []string{"acmeServer"},
		Short:   "Create ACME server",
		Args:    cobra.NoArgs,
	}, func(ctx context.Context, client *apiserver.Client, cmd *cobra.Command, args []string) error {
		var apiACMEServer *apiv1.ACMEServer
		if letsEncrypt {
			apiACMEServer = &apiv1.ACMEServer{
				DirectoryUrl:    "https://acme-v02.api.letsencrypt.org/directory",
				DisplayName:     "Let's Encrypt",
				IntegrationName: "acme",
			}
		} else if letsEncryptStaging {
			apiACMEServer = &apiv1.ACMEServer{
				DirectoryUrl:    "https://acme-staging-v02.api.letsencrypt.org/directory",
				DisplayName:     "Let's Encrypt Staging",
				IntegrationName: "acme",
			}
		} else {
			apiACMEServer = &apiv1.ACMEServer{
				DirectoryUrl:    directoryURL,
				DisplayName:     displayName,
				IntegrationName: integrationName,
			}
		}
		var err error
		if apiACMEServer, err = client.ACMEServer.Create(context.Background(), &apiv1.CreateACMEServerRequest{
			AcmeServer: apiACMEServer,
		}); err != nil {
			return err
		}
		return FormatResource(acmeServer{}.Encode(apiACMEServer), cmd.OutOrStdout())
	})

	cmd.Flags().BoolVar(&letsEncrypt, "letsencrypt", false, "Let's Encrypt defaults")
	cmd.Flags().BoolVar(&letsEncryptStaging, "letsencrypt-staging", false, "Let's Encrypt staging defaults")
	cmd.Flags().StringVar(&directoryURL, "directory-url", "", "Directory URL")
	cmd.Flags().StringVar(&displayName, "display-name", "", "Display name")
	cmd.Flags().StringVar(&integrationName, "integration-name", "", "Integration name")

	return cmd
}

type acmeServer struct{}

func (r acmeServer) Create(ctx context.Context, client *apiserver.Client, resource *Resource) (*Resource, error) {
	spec := resource.Spec.(*acmeServerSpec)
	acmeServer := &apiv1.ACMEServer{
		DirectoryUrl:    spec.DirectoryURL,
		DisplayName:     spec.DisplayName,
		IntegrationName: spec.IntegrationName,
	}
	acmeServer, err := client.ACMEServer.Create(ctx, &apiv1.CreateACMEServerRequest{
		AcmeServer: acmeServer,
	})
	if err != nil {
		return nil, err
	}
	return r.Encode(acmeServer), nil
}

func (r acmeServer) Delete(ctx context.Context, client *apiserver.Client, name string) error {
	_, err := client.ACMEServer.Delete(ctx, &apiv1.DeleteACMEServerRequest{
		Name: fmt.Sprintf("acmeServers/%s", name),
	})
	return err
}

func (r acmeServer) Encode(acmeServer *apiv1.ACMEServer) *Resource {
	uid := strings.Split(acmeServer.Name, "/")[1]
	return &Resource{
		Kind: "acmeserver",
		Meta: &resourceMeta{
			UID:        uid,
			CreateTime: acmeServer.GetCreateTime().AsTime(),
			UpdateTime: acmeServer.GetUpdateTime().AsTime(),
		},
		Spec: &acmeServerSpec{
			DisplayName:     acmeServer.GetDisplayName(),
			DirectoryURL:    acmeServer.GetDirectoryUrl(),
			IntegrationName: acmeServer.GetIntegrationName(),
		},
	}
}

func (r acmeServer) Get(ctx context.Context, client *apiserver.Client, name string) (*Resource, error) {
	acmeServer, err := client.ACMEServer.Get(ctx, &apiv1.GetACMEServerRequest{
		Name: fmt.Sprintf("acmeServers/%s", name),
	})
	if err != nil {
		return nil, err
	}
	return r.Encode(acmeServer), nil
}

func (r acmeServer) List(ctx context.Context, client *apiserver.Client) ([]*Resource, error) {
	return listResource(func(pageToken string) ([]*Resource, string, error) {
		response, err := client.ACMEServer.List(ctx, &apiv1.ListACMEServersRequest{
			PageToken: pageToken,
			PageSize:  0,
		})
		if err != nil {
			return nil, "", err
		}
		a := make([]*Resource, len(response.GetAcmeServers()))
		for i, acmeServer := range response.GetAcmeServers() {
			a[i] = r.Encode(acmeServer)
		}
		return a, response.GetNextPageToken(), nil
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

func (r acmeServer) Describe(_ context.Context, _ *apiserver.Client, resource *Resource, output io.Writer) (err error) {
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

func init() {
	resources.Add(acmeServer{}, "as")
}

package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"powerssl.io/pkg/apiserver/api"
	apiserverclient "powerssl.io/pkg/apiserver/client"
)

var ACMEServer acmeServer

type ACMEServerSpec struct {
	DisplayName     string `json:"displayName,omitempty"     yaml:"displayName,omitempty"`
	DirectoryURL    string `json:"directoryURL,omitempty"    yaml:"directoryURL,omitempty"`
	IntegrationName string `json:"integrationName,omitempty" yaml:"integrationName,omitempty"`
}

type acmeServer struct{}

func (r acmeServer) Create(client *apiserverclient.GRPCClient, resource *Resource) (*Resource, error) {
	spec := resource.Spec.(*ACMEServerSpec)
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

func (r acmeServer) Delete(client *apiserverclient.GRPCClient, name string) error {
	return client.ACMEServer.Delete(context.Background(), fmt.Sprintf("acmeServers/%s", name))
}

func (r acmeServer) Encode(acmeServer *api.ACMEServer) *Resource {
	uid := strings.Split(acmeServer.Name, "/")[1]
	return &Resource{
		Kind: "acmeserver",
		Meta: &ResourceMeta{
			UID:        uid,
			CreateTime: acmeServer.CreateTime,
			UpdateTime: acmeServer.UpdateTime,
		},
		Spec: &ACMEServerSpec{
			DisplayName:     acmeServer.DisplayName,
			DirectoryURL:    acmeServer.DirectoryURL,
			IntegrationName: acmeServer.IntegrationName,
		},
	}
}

func (r acmeServer) Get(client *apiserverclient.GRPCClient, name string) (*Resource, error) {
	acmeServer, err := client.ACMEServer.Get(context.Background(), fmt.Sprintf("acmeServers/%s", name))
	if err != nil {
		return nil, err
	}
	return r.Encode(acmeServer), nil
}

func (r acmeServer) List(client *apiserverclient.GRPCClient) ([]*Resource, error) {
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
	return new(ACMEServerSpec)
}

var (
	DirectoryURL       string
	IntegrationName    string
	LetsEncrypt        bool
	LetsEncryptStaging bool
)

var createACMEServerCmd = &cobra.Command{
	Use:     "acmeserver",
	Aliases: []string{"acmeServer"},
	Short:   "Create ACME server",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := NewGRPCClient()
		if err != nil {
			er(err)
		}
		var acmeServer *api.ACMEServer
		if LetsEncrypt {
			acmeServer = &api.ACMEServer{
				DirectoryURL:    "https://acme-v02.api.letsencrypt.org/directory",
				DisplayName:     "Let's Encrypt",
				IntegrationName: "acme",
			}
		} else if LetsEncryptStaging {
			acmeServer = &api.ACMEServer{
				DirectoryURL:    "https://acme-staging-v02.api.letsencrypt.org/directory",
				DisplayName:     "Let's Encrypt Staging",
				IntegrationName: "acme",
			}
		} else {
			acmeServer = &api.ACMEServer{
				DirectoryURL:    DirectoryURL,
				DisplayName:     DisplayName,
				IntegrationName: IntegrationName,
			}
		}
		acmeServer, err = client.ACMEServer.Create(context.Background(), acmeServer)
		if err != nil {
			er(err)
		}
		pr(ACMEServer.Encode(acmeServer))
	},
}

func init() {
	Resources.Add(ACMEServer, "as")

	createACMEServerCmd.Flags().BoolVarP(&LetsEncrypt, "letsencrypt", "", false, "Let's Encrypt defaults")
	createACMEServerCmd.Flags().BoolVarP(&LetsEncryptStaging, "letsencrypt-staging", "", false, "Let's Encrypt staging defaults")
	createACMEServerCmd.Flags().StringVarP(&DisplayName, "display-name", "", "", "Display name")
	createACMEServerCmd.Flags().StringVarP(&DirectoryURL, "directory-url", "", "", "Directory URL")
	createACMEServerCmd.Flags().StringVarP(&Filename, "filename", "f", "", "Filename to file to use to create the ACME server")
	createACMEServerCmd.Flags().StringVarP(&IntegrationName, "integration-name", "", "", "Integration name")

	createCmd.AddCommand(createACMEServerCmd)
}

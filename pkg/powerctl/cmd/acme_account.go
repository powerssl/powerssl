package cmd

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"powerssl.io/pkg/apiserver/api"
	apiserverclient "powerssl.io/pkg/apiserver/client"
)

var ACMEAccount acmeAccount

type ACMEAccountSpec struct {
	DisplayName          string   `json:"displayName,omitempty"          yaml:"displayName,omitempty"`
	Title                string   `json:"directoryURL,omitempty"         yaml:"directoryURL,omitempty"`
	Description          string   `json:"integrationName,omitempty"      yaml:"integrationName,omitempty"`
	ACMEServer           string   `json:"acmeServer,omitempty"           yaml:"acmeServer,omitempty"`
	TermsOfServiceAgreed bool     `json:"termsOfServiceAgreed,omitempty" yaml:"termsOfServiceAgreed,omitempty"`
	Contacts             []string `json:"contacts,omitempty"             yaml:"contacts,omitempty"`
	AccountURL           string   `json:"accountURL,omitempty"           yaml:"accountURL,omitempty"`
}

type acmeAccount struct{}

func (r acmeAccount) Create(client *apiserverclient.GRPCClient, resource *Resource) (*Resource, error) {
	spec := resource.Spec.(*ACMEAccountSpec)
	acmeAccount := &api.ACMEAccount{
		Contacts:             spec.Contacts,
		TermsOfServiceAgreed: spec.TermsOfServiceAgreed,
	}
	acmeAccount, err := client.ACMEAccount.Create(context.Background(), fmt.Sprintf("acmeServers/%s", spec.ACMEServer), acmeAccount)
	if err != nil {
		return nil, err
	}
	return r.Encode(acmeAccount), nil
}

func (r acmeAccount) Delete(client *apiserverclient.GRPCClient, name string) error {
	return client.ACMEAccount.Delete(context.Background(), fmt.Sprintf("acmeServers/-/acmeAccounts/%s", name))
}

func (r acmeAccount) Encode(acmeAccount *api.ACMEAccount) *Resource {
	uid := strings.Split(acmeAccount.Name, "/")[3]
	acmeServer := strings.Split(acmeAccount.Name, "/")[1]
	return &Resource{
		Kind: "acmeaccount",
		Meta: &ResourceMeta{
			UID:        uid,
			CreateTime: acmeAccount.CreateTime,
			UpdateTime: acmeAccount.UpdateTime,
		},
		Spec: &ACMEAccountSpec{
			DisplayName:          acmeAccount.DisplayName,
			Title:                acmeAccount.Title,
			Description:          acmeAccount.Description,
			ACMEServer:           acmeServer,
			TermsOfServiceAgreed: acmeAccount.TermsOfServiceAgreed,
			Contacts:             acmeAccount.Contacts,
			AccountURL:           acmeAccount.AccountURL,
		},
	}
}

func (r acmeAccount) Get(client *apiserverclient.GRPCClient, name string) (*Resource, error) {
	acmeAccount, err := client.ACMEAccount.Get(context.Background(), fmt.Sprintf("acmeServers/-/acmeAccounts/%s", name))
	if err != nil {
		return nil, err
	}
	return r.Encode(acmeAccount), nil
}

func (r acmeAccount) List(client *apiserverclient.GRPCClient) ([]*Resource, error) {
	return listResource(func(pageToken string) ([]*Resource, string, error) {
		acmeAccounts, nextPageToken, err := client.ACMEAccount.List(context.Background(), "parent", 0, pageToken)
		if err != nil {
			return nil, nextPageToken, err
		}
		a := make([]*Resource, len(acmeAccounts))
		for i, acmeAccount := range acmeAccounts {
			a[i] = r.Encode(acmeAccount)
		}
		return a, nextPageToken, nil
	})
}

func (r acmeAccount) Spec() interface{} {
	return new(ACMEAccountSpec)
}

func (r acmeAccount) Columns(resource *Resource) ([]string, []string) {
	spec := resource.Spec.(*ACMEAccountSpec)
	return []string{
			"DISPLAY NAME",
			"DESCRIPTION",
			"ACME SERVER",
			"TOS AGREED",
			"CONTACTS",
		}, []string{
			fmt.Sprint(spec.DisplayName),
			fmt.Sprint(spec.Description),
			fmt.Sprint(spec.ACMEServer),
			fmt.Sprint(spec.TermsOfServiceAgreed),
			strings.Join(spec.Contacts, ","),
		}
}

func (r acmeAccount) Describe(client *apiserverclient.GRPCClient, resource *Resource, output io.Writer) (err error) {
	spec := resource.Spec.(*ACMEAccountSpec)
	w := tabwriter.NewWriter(output, 0, 0, 1, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, fmt.Sprintf("UID:\t%s", resource.Meta.UID))
	fmt.Fprintln(w, fmt.Sprintf("Create Time:\t%s", resource.Meta.CreateTime))
	fmt.Fprintln(w, fmt.Sprintf("Update Time:\t%s", resource.Meta.UpdateTime))
	fmt.Fprintln(w, fmt.Sprintf("Display Name:\t%s", spec.DisplayName))
	fmt.Fprintln(w, fmt.Sprintf("Title:\t%s", spec.Title))
	fmt.Fprintln(w, fmt.Sprintf("Description:\t%s", spec.Description))
	fmt.Fprintln(w, fmt.Sprintf("TOS Agreed:\t%v", spec.TermsOfServiceAgreed))
	fmt.Fprintln(w, fmt.Sprintf("Contacts:\t%s", strings.Join(spec.Contacts, ",")))
	fmt.Fprintln(w, fmt.Sprintf("Account URL:\t%s", spec.AccountURL))
	fmt.Fprintln(w, "ACME Server:")
	acmeServer := &Resource{
		Kind: "acmeserver",
		Meta: &ResourceMeta{
			UID: spec.ACMEServer,
		},
	}
	if acmeServer, err = acmeServer.Get(client); err != nil {
		return err
	}
	acmeServerDescription := new(bytes.Buffer)
	acmeServer.Describe(client, acmeServerDescription)
	scanner := bufio.NewScanner(acmeServerDescription)
	for scanner.Scan() {
		fmt.Fprintln(w, "", "", scanner.Text())
	}
	w.Flush()
	return nil
}

var (
	ACMEServerID         string
	Contacts             string
	TermsOfServiceAgreed bool
)

var createACMEAccountCmd = &cobra.Command{
	Use:     "acmeaccount",
	Short:   "Create ACME account",
	Args:    cobra.NoArgs,
	Example: `  powerctl create acmeaccount --agree-terms-of-service --contacts mailto:john.doe@example.com --acmeserver 42   Create ACME account within ACME server`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := NewGRPCClient()
		if err != nil {
			er(err)
		}
		acmeAccount := &api.ACMEAccount{
			Contacts:             strings.Split(Contacts, ","),
			TermsOfServiceAgreed: TermsOfServiceAgreed,
		}
		acmeAccount, err = client.ACMEAccount.Create(context.Background(), fmt.Sprintf("acmeServers/%s", ACMEServerID), acmeAccount)
		if err != nil {
			er(err)
		}
		pr(ACMEAccount.Encode(acmeAccount))
	},
}

func init() {
	Resources.Add(ACMEAccount, "aa")

	createACMEAccountCmd.Flags().BoolVarP(&TermsOfServiceAgreed, "agree-terms-of-service", "", false, "Terms of Service agreed")
	createACMEAccountCmd.Flags().StringVarP(&Contacts, "contacts", "", "", "Contact URLs (e.g. mailto:contact@example.com) (seperated by \",\")")
	createACMEAccountCmd.Flags().StringVarP(&Filename, "filename", "f", "", "Filename to file to use to create the ACME account")
	createACMEAccountCmd.Flags().StringVarP(&ACMEServerID, "acmeserver", "", "", "ACME Server")
	createACMEAccountCmd.MarkFlagRequired("acmeserver")

	createCmd.AddCommand(createACMEAccountCmd)
}

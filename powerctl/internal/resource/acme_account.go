package resource

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"

	cmdutil "powerssl.dev/common/cmd"
	"powerssl.dev/powerctl/internal"
	"powerssl.dev/sdk/apiserver/api"
	apiserverclient "powerssl.dev/sdk/apiserver/client"
)

type acmeAccount struct{}

func (r acmeAccount) Create(client *apiserverclient.GRPCClient, resource *Resource) (*Resource, error) {
	spec := resource.Spec.(*acmeAccountSpec)
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
		Meta: &resourceMeta{
			UID:        uid,
			CreateTime: acmeAccount.CreateTime,
			UpdateTime: acmeAccount.UpdateTime,
		},
		Spec: &acmeAccountSpec{
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
	return new(acmeAccountSpec)
}

func (r acmeAccount) Columns(resource *Resource) ([]string, []string) {
	spec := resource.Spec.(*acmeAccountSpec)
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
	spec := resource.Spec.(*acmeAccountSpec)
	w := tabwriter.NewWriter(output, 0, 0, 1, ' ', tabwriter.TabIndent)
	_, _ = fmt.Fprintln(w, fmt.Sprintf("UID:\t%s", resource.Meta.UID))
	_, _ = fmt.Fprintln(w, fmt.Sprintf("Create Time:\t%s", resource.Meta.CreateTime))
	_, _ = fmt.Fprintln(w, fmt.Sprintf("Update Time:\t%s", resource.Meta.UpdateTime))
	_, _ = fmt.Fprintln(w, fmt.Sprintf("Display Name:\t%s", spec.DisplayName))
	_, _ = fmt.Fprintln(w, fmt.Sprintf("Title:\t%s", spec.Title))
	_, _ = fmt.Fprintln(w, fmt.Sprintf("Description:\t%s", spec.Description))
	_, _ = fmt.Fprintln(w, fmt.Sprintf("TOS Agreed:\t%v", spec.TermsOfServiceAgreed))
	_, _ = fmt.Fprintln(w, fmt.Sprintf("Contacts:\t%s", strings.Join(spec.Contacts, ",")))
	_, _ = fmt.Fprintln(w, fmt.Sprintf("Account URL:\t%s", spec.AccountURL))
	_, _ = fmt.Fprintln(w, "ACME Server:")
	acmeServer := &Resource{
		Kind: "acmeserver",
		Meta: &resourceMeta{
			UID: spec.ACMEServer,
		},
	}
	if acmeServer, err = acmeServer.Get(client); err != nil {
		return err
	}
	acmeServerDescription := new(bytes.Buffer)
	if err = acmeServer.Describe(client, acmeServerDescription); err != nil {
		return err
	}
	scanner := bufio.NewScanner(acmeServerDescription)
	for scanner.Scan() {
		_, _ = fmt.Fprintln(w, "", "", scanner.Text())
	}
	return w.Flush()
}

type acmeAccountSpec struct {
	DisplayName          string   `json:"displayName,omitempty"          yaml:"displayName,omitempty"`
	Title                string   `json:"directoryURL,omitempty"         yaml:"directoryURL,omitempty"`
	Description          string   `json:"integrationName,omitempty"      yaml:"integrationName,omitempty"`
	ACMEServer           string   `json:"acmeServer,omitempty"           yaml:"acmeServer,omitempty"`
	TermsOfServiceAgreed bool     `json:"termsOfServiceAgreed,omitempty" yaml:"termsOfServiceAgreed,omitempty"`
	Contacts             []string `json:"contacts,omitempty"             yaml:"contacts,omitempty"`
	AccountURL           string   `json:"accountURL,omitempty"           yaml:"accountURL,omitempty"`
}

func NewCmdCreateACMEAccount() *cobra.Command {
	var client *apiserverclient.GRPCClient
	var (
		acmeServerID         string
		contacts             string
		termsOfServiceAgreed bool
	)

	cmd := &cobra.Command{
		Use:     "acmeaccount",
		Short:   "Create ACME account",
		Args:    cobra.NoArgs,
		Example: `  powerctl create acmeaccount --agree-terms-of-service --contacts mailto:john.doe@example.com --acmeserver 42   Create ACME account within ACME server`,
		PreRunE: func(cmd *cobra.Command, args []string) (err error) {
			client, err = internal.NewGRPCClient()
			return err
		},
		Run: cmdutil.HandleError(func(cmd *cobra.Command, args []string) (err error) {
			apiACMEAccount := &api.ACMEAccount{
				Contacts:             strings.Split(contacts, ","),
				TermsOfServiceAgreed: termsOfServiceAgreed,
			}
			if apiACMEAccount, err = client.ACMEAccount.Create(context.Background(), fmt.Sprintf("acmeServers/%s", acmeServerID), apiACMEAccount); err != nil {
				return err
			}
			return FormatResource(acmeAccount{}.Encode(apiACMEAccount), cmd.OutOrStdout())
		}),
	}

	cmd.Flags().BoolVar(&termsOfServiceAgreed, "agree-terms-of-service", false, "Terms of Service agreed")
	cmd.Flags().StringVar(&contacts, "contacts", "", "Contact URLs (e.g. mailto:contact@example.com) (seperated by \",\")")
	cmd.Flags().StringVar(&acmeServerID, "acmeserver", "", "ACME Server")

	cmdutil.Must(cmd.MarkFlagRequired("acmeserver"))

	return cmd
}

func init() {
	resources.Add(acmeAccount{}, "aa")
}

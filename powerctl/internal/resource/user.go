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
	"powerssl.dev/powerctl/internal"
	"powerssl.dev/sdk/apiserver"
	"powerssl.dev/sdk/apiserver/api"
)

type user struct{}

func (r user) Create(client *apiserver.Client, resource *Resource) (*Resource, error) {
	spec := resource.Spec.(*userSpec)
	user := &api.User{
		DisplayName: spec.DisplayName,
		UserName:    spec.UserName,
	}
	user, err := client.User.Create(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return r.Encode(user), nil
}

func (r user) Delete(client *apiserver.Client, name string) error {
	return client.User.Delete(context.Background(), fmt.Sprintf("users/%s", name))
}

func (r user) Encode(user *api.User) *Resource {
	uid := strings.Split(user.Name, "/")[1]
	return &Resource{
		Kind: "user",
		Meta: &resourceMeta{
			UID:        uid,
			CreateTime: user.CreateTime,
			UpdateTime: user.UpdateTime,
		},
		Spec: &userSpec{
			DisplayName: user.DisplayName,
			UserName:    user.UserName,
		},
	}
}

func (r user) Get(client *apiserver.Client, name string) (*Resource, error) {
	user, err := client.User.Get(context.Background(), fmt.Sprintf("users/%s", name))
	if err != nil {
		return nil, err
	}
	return r.Encode(user), nil
}

func (r user) List(client *apiserver.Client) ([]*Resource, error) {
	return listResource(func(pageToken string) ([]*Resource, string, error) {
		users, nextPageToken, err := client.User.List(context.Background(), 0, pageToken)
		if err != nil {
			return nil, nextPageToken, err
		}
		a := make([]*Resource, len(users))
		for i, user := range users {
			a[i] = r.Encode(user)
		}
		return a, nextPageToken, nil
	})
}

func (r user) Spec() interface{} {
	return new(userSpec)
}

func (r user) Columns(resource *Resource) ([]string, []string) {
	spec := resource.Spec.(*userSpec)
	return []string{
			"DISPLAY NAME",
			"USER NAME",
		}, []string{
			fmt.Sprint(spec.DisplayName),
			fmt.Sprint(spec.UserName),
		}
}

func (r user) Describe(_ *apiserver.Client, resource *Resource, output io.Writer) (err error) {
	spec := resource.Spec.(*userSpec)
	w := tabwriter.NewWriter(output, 0, 0, 1, ' ', tabwriter.TabIndent)
	_, _ = fmt.Fprintln(w, fmt.Sprintf("UID:\t%s", resource.Meta.UID))
	_, _ = fmt.Fprintln(w, fmt.Sprintf("Create Time:\t%s", resource.Meta.CreateTime))
	_, _ = fmt.Fprintln(w, fmt.Sprintf("Update Time:\t%s", resource.Meta.UpdateTime))
	_, _ = fmt.Fprintln(w, fmt.Sprintf("Display Name:\t%s", spec.DisplayName))
	_, _ = fmt.Fprintln(w, fmt.Sprintf("User Name:\t%s", spec.UserName))
	return w.Flush()
}

type userSpec struct {
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	UserName    string `json:"userName,omitempty"    yaml:"userName,omitempty"`
}

func NewCmdCreateUser() *cobra.Command {
	var client *apiserver.Client
	var (
		displayName string
		userName    string
	)

	cmd := &cobra.Command{
		Use:     "user",
		Aliases: []string{"user"},
		Short:   "Create ACME server",
		Args:    cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) (err error) {
			client, err = internal.NewGRPCClient()
			return err
		},
		Run: cmdutil.HandleError(func(cmd *cobra.Command, args []string) (err error) {
			apiUser := &api.User{
				DisplayName: displayName,
				UserName:    userName,
			}
			if apiUser, err = client.User.Create(context.Background(), apiUser); err != nil {
				return err
			}
			return FormatResource(user{}.Encode(apiUser), os.Stdout)
		}),
	}

	cmd.Flags().StringVar(&displayName, "display-name", "", "Display name")
	cmd.Flags().StringVar(&userName, "user-name", "", "User name")

	return cmd
}

func init() {
	resources.Add(user{}, "u")
}

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

func NewCmdCreateUser() *cobra.Command {
	var (
		displayName string
		userName    string
	)

	cmd := internal.CmdWithClient(&cobra.Command{
		Use:     "user",
		Aliases: []string{"user"},
		Short:   "Create ACME server",
		Args:    cobra.NoArgs,
	}, func(ctx context.Context, client *apiserver.Client, cmd *cobra.Command, args []string) error {
		apiUser := &apiv1.User{
			DisplayName: displayName,
			UserName:    userName,
		}
		var err error
		if apiUser, err = client.User.Create(context.Background(), &apiv1.CreateUserRequest{
			User: apiUser,
		}); err != nil {
			return err
		}
		return FormatResource(user{}.Encode(apiUser), cmd.OutOrStdout())
	})

	cmd.Flags().StringVar(&displayName, "display-name", "", "Display name")
	cmd.Flags().StringVar(&userName, "user-name", "", "User name")

	return cmd
}

type user struct{}

func (r user) Create(ctx context.Context, client *apiserver.Client, resource *Resource) (*Resource, error) {
	spec := resource.Spec.(*userSpec)
	user := &apiv1.User{
		DisplayName: spec.DisplayName,
		UserName:    spec.UserName,
	}
	user, err := client.User.Create(ctx, &apiv1.CreateUserRequest{
		User: user,
	})
	if err != nil {
		return nil, err
	}
	return r.Encode(user), nil
}

func (r user) Delete(ctx context.Context, client *apiserver.Client, name string) error {
	_, err := client.User.Delete(ctx, &apiv1.DeleteUserRequest{
		Name: fmt.Sprintf("users/%s", name),
	})
	return err
}

func (r user) Encode(user *apiv1.User) *Resource {
	uid := strings.Split(user.GetName(), "/")[1]
	return &Resource{
		Kind: "user",
		Meta: &resourceMeta{
			UID:        uid,
			CreateTime: user.GetCreateTime().AsTime(),
			UpdateTime: user.GetUpdateTime().AsTime(),
		},
		Spec: &userSpec{
			DisplayName: user.DisplayName,
			UserName:    user.UserName,
		},
	}
}

func (r user) Get(ctx context.Context, client *apiserver.Client, name string) (*Resource, error) {
	user, err := client.User.Get(ctx, &apiv1.GetUserRequest{
		Name: fmt.Sprintf("users/%s", name),
	})
	if err != nil {
		return nil, err
	}
	return r.Encode(user), nil
}

func (r user) List(ctx context.Context, client *apiserver.Client) ([]*Resource, error) {
	return listResource(func(pageToken string) ([]*Resource, string, error) {
		response, err := client.User.List(ctx, &apiv1.ListUsersRequest{
			PageSize:  0,
			PageToken: pageToken,
		})
		if err != nil {
			return nil, "", err
		}
		a := make([]*Resource, len(response.GetUsers()))
		for i, user := range response.GetUsers() {
			a[i] = r.Encode(user)
		}
		return a, response.GetNextPageToken(), nil
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

func (r user) Describe(_ context.Context, _ *apiserver.Client, resource *Resource, output io.Writer) (err error) {
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

func init() {
	resources.Add(user{}, "u")
}

package model

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"

	"powerssl.dev/powerssl/internal/pkg/uid"
	"powerssl.dev/powerssl/pkg/apiserver/api"
)

type User struct {
	ID        string `pg:",pk"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `pg:",soft_delete"`

	DisplayName string
	UserName    string
}

func (u *User) Name() string {
	return fmt.Sprintf("users/%s", u.ID)
}

var _ pg.BeforeInsertHook = (*User)(nil)

func (user *User) BeforeInsert(ctx context.Context) (context.Context, error) {
	user.ID = uid.New()
	return ctx, nil
}

func (u *User) ToAPI() *api.User {
	return &api.User{
		Name: u.Name(),

		CreateTime:  u.CreatedAt,
		UpdateTime:  u.UpdatedAt,
		DisplayName: u.DisplayName,

		UserName: u.UserName,
	}
}

type Users []*User

func (a Users) ToAPI() []*api.User {
	servers := make([]*api.User, len(a))
	for i, server := range a {
		servers[i] = server.ToAPI()
	}
	return servers
}

func FindUserByName(name string, db *pg.DB) (*User, error) {
	s := strings.Split(name, "/")
	if len(s) != 2 {
		return nil, status.Error(codes.InvalidArgument, "malformed name")
	}

	user := &User{}
	if err := db.Model(user).Where("id = ?", s[1]).Limit(1).Select(); err != nil {
		if err == pg.ErrNoRows {
			return nil, status.Error(codes.NotFound, "not found")
		}
		return nil, err
	}
	return user, nil
}

func NewUserFromAPI(user *api.User) *User {
	return &User{
		DisplayName: user.DisplayName,
		UserName:    user.UserName,
	}
}

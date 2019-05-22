package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/gogo/status"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"

	"powerssl.io/powerssl/internal/pkg/uid"
	"powerssl.io/powerssl/pkg/apiserver/api"
)

type User struct {
	ID        string `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	DisplayName string
	UserName    string
}

func (*User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uid.New())
	return nil
}

func (u *User) Name() string {
	return fmt.Sprintf("users/%s", u.ID)
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

func FindUserByName(name string, db *gorm.DB) (*User, error) {
	s := strings.Split(name, "/")
	if len(s) != 2 {
		return nil, status.Error(codes.InvalidArgument, "malformed name")
	}

	user := &User{}
	if db.Where("id = ?", s[1]).First(&user).RecordNotFound() {
		return nil, status.Error(codes.NotFound, "not found")
	}
	return user, nil
}

func NewUserFromAPI(user *api.User) *User {
	return &User{
		DisplayName: user.DisplayName,
		UserName:    user.UserName,
	}
}

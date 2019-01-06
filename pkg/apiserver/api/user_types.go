package api // import "powerssl.io/pkg/apiserver/api"

import "time"

type User struct {
	Name        string    `json:"name,omitempty"        yaml:"name,omitempty"`
	CreateTime  time.Time `json:"createTime,omitempty"  yaml:"createTime,omitempty"`
	UpdateTime  time.Time `json:"updateTime,omitempty"  yaml:"updateTime,omitempty"`
	DisplayName string    `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	UserName    string    `json:"userName,omitempty"    yaml:"userName,omitempty"`
}

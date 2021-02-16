package api // import "powerssl.dev/sdk/apiserver/api"

import "time"

type ACMEAccount struct {
	Name                 string            `json:"name,omitempty"                 yaml:"name,omitempty"`
	CreateTime           time.Time         `json:"createTime,omitempty"           yaml:"createTime,omitempty"`
	UpdateTime           time.Time         `json:"updateTime,omitempty"           yaml:"updateTime,omitempty"`
	DisplayName          string            `json:"displayName,omitempty"          yaml:"displayName,omitempty"`
	Title                string            `json:"title,omitempty"                yaml:"title,omitempty"`
	Description          string            `json:"description,omitempty"          yaml:"description,omitempty"`
	Labels               map[string]string `json:"labels,omitempty"               yaml:"labels,omitempty"`
	TermsOfServiceAgreed bool              `json:"termsOfServiceAgreed,omitempty" yaml:"termsOfServiceAgreed,omitempty"`
	Contacts             []string          `json:"contacts,omitempty"             yaml:"contacts,omitempty"`
	AccountURL           string            `json:"accountURL,omitempty"           yaml:"accountURL,omitempty"`
}

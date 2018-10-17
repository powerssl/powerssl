package api // import "powerssl.io/pkg/apiserver/api"

import "time"

type Certificate struct {
	Name            string            `json:"name,omitempty"            yaml:"name,omitempty"`
	CreateTime      time.Time         `json:"createTime,omitempty"      yaml:"createTime,omitempty"`
	UpdateTime      time.Time         `json:"updateTime,omitempty"      yaml:"updateTime,omitempty"`
	DisplayName     string            `json:"displayName,omitempty"     yaml:"displayName,omitempty"`
	Title           string            `json:"title,omitempty"           yaml:"title,omitempty"`
	Description     string            `json:"description,omitempty"     yaml:"description,omitempty"`
	Labels          map[string]string `json:"labels,omitempty"          yaml:"labels,omitempty"`
	Dnsnames        []string          `json:"dnsnames,omitempty"        yaml:"dnsnames,omitempty"`
	KeyAlgorithm    string            `json:"keyAlgorithm,omitempty"    yaml:"keyAlgorithm,omitempty"`
	KeySize         int32             `json:"keySize,omitempty"         yaml:"keySize,omitempty"`
	DigestAlgorithm string            `json:"digestAlgorithm,omitempty" yaml:"digestAlgorithm,omitempty"`
	AutoRenew       bool              `json:"autoRenew,omitempty"       yaml:"autoRenew,omitempty"`
}

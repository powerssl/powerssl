package api // import "powerssl.dev/sdk/apiserver/api"

import "time"

type ACMEServer struct {
	Name            string    `json:"name,omitempty"            yaml:"name,omitempty"`
	CreateTime      time.Time `json:"createTime,omitempty"      yaml:"createTime,omitempty"`
	UpdateTime      time.Time `json:"updateTime,omitempty"      yaml:"updateTime,omitempty"`
	DisplayName     string    `json:"displayName,omitempty"     yaml:"displayName,omitempty"`
	DirectoryURL    string    `json:"directoryURL,omitempty"    yaml:"directoryURL,omitempty"`
	IntegrationName string    `json:"integrationName,omitempty" yaml:"integrationName,omitempty"`
}

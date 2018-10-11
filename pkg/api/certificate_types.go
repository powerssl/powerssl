package api // import "powerssl.io/pkg/api"

import "time"

type Certificate struct {
	Name            string
	CreateTime      time.Time
	UpdateTime      time.Time
	DisplayName     string
	Title           string
	Description     string
	Labels          map[string]string
	Dnsnames        []string
	KeyAlgorithm    string
	KeySize         int32
	DigestAlgorithm string
	AutoRenew       bool
}

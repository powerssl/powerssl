package cloudflare

import (
	integrationdns "powerssl.io/pkg/integration/dns"
)

type errorConst string

func (e errorConst) Error() string { return string(e) }

const ErrNotImplemented = errorConst("not implemented.")

type Cloudflare struct{}

func New() integrationdns.Integration {
	return &Cloudflare{}
}

func (cloudflare *Cloudflare) CreateRecord(domain, recordType, content string) (err error) {
	return err
}

func (cloudflare *Cloudflare) DeleteRecord(domain, recordType string) (err error) {
	return err
}

func (cloudflare *Cloudflare) VerifyDomain(domain string) (err error) {
	return err
}

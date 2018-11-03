package cloudflare

import (
	"context"

	integrationdns "powerssl.io/pkg/integration/dns"
)

type errorConst string

func (e errorConst) Error() string { return string(e) }

const ErrNotImplemented = errorConst("not implemented.")

type Cloudflare struct{}

func New() integrationdns.Integration {
	return &Cloudflare{}
}

func (cloudflare *Cloudflare) CreateRecord(_ context.Context, domain, recordType, content string) (err error) {
	return err
}

func (cloudflare *Cloudflare) DeleteRecord(_ context.Context, domain, recordType string) (err error) {
	return err
}

func (cloudflare *Cloudflare) VerifyDomain(_ context.Context, domain string) (err error) {
	return err
}

package cloudflare

import "context"

type cloudflare struct{}

func New() *cloudflare {
	return &cloudflare{}
}

func (cloudflare *cloudflare) CreateRecord(ctx context.Context, domain, recordType, content string) (err error) {
	return err
}

func (cloudflare *cloudflare) DeleteRecord(ctx context.Context, domain, recordType string) (err error) {
	return err
}

func (cloudflare *cloudflare) VerifyDomain(cxt context.Context, domain string) (err error) {
	return err
}

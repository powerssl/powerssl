package cloudflare

import (
	"powerssl.io/pkg/integration"
)

const Name = "Cloudflare"

type Cloudflare struct{}

func New() integration.DNSIntegration {
	return &Cloudflare{}
}

func (cloudflare *Cloudflare) GetName() string {
	return Name
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

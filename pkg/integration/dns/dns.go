package dns

import (
	"powerssl.io/pkg/controller/api"
)

type Integration interface {
	CreateRecord(domain string, recordType string, content string) (err error)
	DeleteRecord(domain string, recordType string) (err error)
	VerifyDomain(domain string) (err error)
}

type integration struct {
	handler Integration
}

func (i *integration) dnsCreateRecord(activity *api.Activity) error {
	return i.handler.CreateRecord("domain", "recordType", "content")
}

func (i *integration) dnsDeleteRecord(activity *api.Activity) error {
	return i.handler.DeleteRecord("domain", "recordType")
}

func (i *integration) dnsVerifyDomain(activity *api.Activity) error {
	return i.handler.VerifyDomain("domain")
}

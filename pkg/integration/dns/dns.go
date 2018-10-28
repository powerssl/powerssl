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
	client  interface{} // TODO
	handler Integration
}

func New(client interface{}, handler Integration) *integration {
	return &integration{
		client:  client,
		handler: handler,
	}
}

func (i *integration) dnsCreateRecord(activity *api.Activity) error {
	i.handler.CreateRecord("domain", "recordType", "content")

	return nil
}

func (i *integration) dnsDeleteRecord(activity *api.Activity) error {
	i.handler.DeleteRecord("domain", "recordType")

	return nil
}

func (i *integration) dnsVerifyDomain(activity *api.Activity) error {
	i.handler.VerifyDomain("domain")

	return nil
}

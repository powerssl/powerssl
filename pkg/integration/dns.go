package integration

import (
	"powerssl.io/pkg/controller/api"
)

type DNSIntegration interface {
	GetName() string

	CreateRecord(domain string, recordType string, content string) (err error)
	DeleteRecord(domain string, recordType string) (err error)
	VerifyDomain(domain string) (err error)
}

func (i *integration) dnsCreateRecord(activity *api.Activity) error {
	handler, ok := i.handler.(DNSIntegration)
	if !ok {
		return errorNotImplemented
	}
	err := handler.CreateRecord("domain", "recordType", "content")
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	var _ = erro
	return nil
}

func (i *integration) dnsDeleteRecord(activity *api.Activity) error {
	handler, ok := i.handler.(DNSIntegration)
	if !ok {
		return errorNotImplemented
	}
	err := handler.DeleteRecord("domain", "recordType")
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	var _ = erro
	return nil
}

func (i *integration) dnsVerifyDomain(activity *api.Activity) error {
	handler, ok := i.handler.(DNSIntegration)
	if !ok {
		return errorNotImplemented
	}
	err := handler.VerifyDomain("domain")
	var erro *api.Error
	if err != nil {
		erro = &api.Error{Message: err.Error()}
	}
	var _ = erro
	return nil
}

package dns

import (
	"context"

	"powerssl.io/pkg/controller/api"
)

type Integration interface {
	CreateRecord(ctx context.Context, domain string, recordType string, content string) (err error)
	DeleteRecord(ctx context.Context, domain string, recordType string) (err error)
	VerifyDomain(ctx context.Context, domain string) (err error)
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

func (i *integration) dnsCreateRecord(ctx context.Context, activity *api.Activity) error {
	i.handler.CreateRecord(ctx, "domain", "recordType", "content")

	return nil
}

func (i *integration) dnsDeleteRecord(ctx context.Context, activity *api.Activity) error {
	i.handler.DeleteRecord(ctx, "domain", "recordType")

	return nil
}

func (i *integration) dnsVerifyDomain(ctx context.Context, activity *api.Activity) error {
	i.handler.VerifyDomain(ctx, "domain")

	return nil
}

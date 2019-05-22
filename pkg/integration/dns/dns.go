package dns

import (
	"context"
	"fmt"

	"powerssl.io/powerssl/pkg/controller/api"
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

func (i *integration) HandleActivity(ctx context.Context, activity *api.Activity) error {
	var err error
	switch activity.Name {
	case api.Activity_DNS_CREATE_RECORD:
		err = i.createRecord(ctx, activity)
	case api.Activity_DNS_DELETE_RECORD:
		err = i.deleteRecord(ctx, activity)
	case api.Activity_DNS_VERIFY_DOMAIN:
		err = i.verifyDomain(ctx, activity)
	default:
		err = fmt.Errorf("activity %s not implemented", activity.Name)
	}
	return err
}

func (i *integration) createRecord(ctx context.Context, activity *api.Activity) error {
	i.handler.CreateRecord(ctx, "domain", "recordType", "content")

	return nil
}

func (i *integration) deleteRecord(ctx context.Context, activity *api.Activity) error {
	i.handler.DeleteRecord(ctx, "domain", "recordType")

	return nil
}

func (i *integration) verifyDomain(ctx context.Context, activity *api.Activity) error {
	i.handler.VerifyDomain(ctx, "domain")

	return nil
}

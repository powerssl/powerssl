package dns // import "powerssl.dev/sdk/integration/dns"

import (
	"context"
	"fmt"

	apiv1 "powerssl.dev/api/controller/v1"
)

type Integration interface {
	CreateRecord(ctx context.Context, domain string, recordType string, content string) error
	DeleteRecord(ctx context.Context, domain string, recordType string) error
	VerifyDomain(ctx context.Context, domain string) error
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

func (i *integration) HandleActivity(ctx context.Context, activity *apiv1.Activity) error {
	switch activity.GetName() {
	case apiv1.Activity_DNS_CREATE_RECORD:
		return i.createRecord(ctx, activity)
	case apiv1.Activity_DNS_DELETE_RECORD:
		return i.deleteRecord(ctx, activity)
	case apiv1.Activity_DNS_VERIFY_DOMAIN:
		return i.verifyDomain(ctx, activity)
	default:
		return fmt.Errorf("activity %s not implemented", activity.Name)
	}
}

func (i *integration) createRecord(ctx context.Context, activity *apiv1.Activity) error {
	_ = i.handler.CreateRecord(ctx, "domain", "recordType", "content")

	return nil
}

func (i *integration) deleteRecord(ctx context.Context, activity *apiv1.Activity) error {
	_ = i.handler.DeleteRecord(ctx, "domain", "recordType")

	return nil
}

func (i *integration) verifyDomain(ctx context.Context, activity *apiv1.Activity) error {
	_ = i.handler.VerifyDomain(ctx, "domain")

	return nil
}

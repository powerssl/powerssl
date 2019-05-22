package workflow

import (
	"context"

	"powerssl.io/powerssl/internal/app/controller/workflow/engine/activity"
	"powerssl.io/powerssl/internal/app/controller/workflow/engine/activity/acme"
	"powerssl.io/powerssl/pkg/controller/api"
)

type RequestCertificate struct {
	DirectoryURL string
	AccountURL   string
	Dnsnames     []string
	NotBefore    string
	NotAfter     string
}

func (w RequestCertificate) Kind() string {
	return "RequestCertificate"
}

func (w RequestCertificate) Run(ctx context.Context) {
	identifiers := make([]*api.Identifier, len(w.Dnsnames))
	for i, dnsname := range w.Dnsnames {
		identifiers[i] = &api.Identifier{Type: api.IdentifierTypeDNS, Value: dnsname}
	}

	{
		a := activity.New(api.Activity_ACME_CREATE_ORDER)
		a.SetInput(acme.CreateOrderInput{w.DirectoryURL, w.AccountURL, identifiers, w.NotBefore, w.NotAfter})
		a.Execute(ctx)
	}
	for range []struct{}{} {
		activity.New(api.Activity_ACME_GET_AUTHORIZATION)
		activity.New(api.Activity_ACME_VALIDATE_CHALLENGE)
	}
	activity.New(api.Activity_ACME_FINALIZE_ORDER)
	activity.New(api.Activity_ACME_GET_CERTIFICATE)
}

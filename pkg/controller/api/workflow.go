package api

type WorkflowKind int32

const (
	WorkflowKindCreateACMEAccount      WorkflowKind = 1
	WorkflowKindRequestACMECertificate WorkflowKind = 2
)

type WorkflowIntegrationFilter struct {
	Kind IntegrationKind
	Name string
}

type WorkflowInput interface {
	WorkflowInput()
}

type CreateACMEAccountInput struct {
	DirectoryURL         string
	TermsOfServiceAgreed bool
	Contacts             []string
}

func (*CreateACMEAccountInput) WorkflowInput() {}

type RequestACMECertificateInput struct {
	DirectoryURL string
	AccountURL   string
	Dnsnames     []string
	NotBefore    string
	NotAfter     string
}

func (*RequestACMECertificateInput) WorkflowInput() {}

type Workflow struct {
	Name               string
	Kind               WorkflowKind
	IntegrationFilters []*WorkflowIntegrationFilter
	Input              WorkflowInput
}

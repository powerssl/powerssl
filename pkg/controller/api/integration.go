package api // import "powerssl.io/pkg/controller/api"

type IntegrationKind int32

const (
	IntegrationKindACME IntegrationKind = 1
	IntegrationKindDNS  IntegrationKind = 2
)

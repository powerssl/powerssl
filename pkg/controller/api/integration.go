package api // import "powerssl.dev/powerssl/pkg/controller/api"

type IntegrationKind int32

const (
	IntegrationKindACME IntegrationKind = 1
	IntegrationKindDNS  IntegrationKind = 2
)

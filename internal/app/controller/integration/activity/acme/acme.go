package acme

import "powerssl.dev/powerssl/pkg/controller/api"

type CreateOrderInput struct {
	DirectoryURL string
	AccountURL   string
	Identifiers  []*api.Identifier
	NotBefore    string
	NotAfter     string
}

type CreateOrderResult struct {
	Order *api.Order
	Error *api.Error
}

type CreateOrder struct {
	Input  CreateOrderInput
	Result CreateOrderResult
}

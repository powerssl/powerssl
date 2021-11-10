package acme

import (
	apiv1 "powerssl.dev/api/controller/v1"
)

type CreateOrderInput struct {
	DirectoryURL string
	AccountURL   string
	Identifiers  []*apiv1.Identifier
	NotBefore    string
	NotAfter     string
}

type CreateOrderResult struct {
	Order *apiv1.Order
	Error *apiv1.Error
}

type CreateOrder struct {
	Input  CreateOrderInput
	Result CreateOrderResult
}

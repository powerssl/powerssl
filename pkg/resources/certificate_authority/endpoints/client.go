package endpoints

import (
	"context"

	"powerssl.io/pkg/api"
)

func (e Endpoints) Create(ctx context.Context, ca *api.CertificateAuthority) (*api.CertificateAuthority, error) {
	resp, err := e.CreateEndpoint(ctx, CreateRequest{
		CertificateAuthority: ca,
	})
	if err != nil {
		return nil, err
	}
	response := resp.(CreateResponse)
	return response.CertificateAuthority, nil
}
func (e Endpoints) Delete(ctx context.Context, name string) error {
	// TODO
	return nil
}
func (e Endpoints) Get(ctx context.Context, name string) (*api.CertificateAuthority, error) {
	// TODO
	return nil, nil
}
func (e Endpoints) List(ctx context.Context) ([]*api.CertificateAuthority, error) {
	// TODO
	return []*api.CertificateAuthority{}, nil
}
func (e Endpoints) Update(ctx context.Context, ca *api.CertificateAuthority) (*api.CertificateAuthority, error) {
	// TODO
	return ca, nil
}

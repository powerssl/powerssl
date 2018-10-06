package endpoints

import (
	"context"

	"powerssl.io/pkg/api"
)

func (e Endpoints) Create(ctx context.Context, ca *api.CertificateAuthority) (*api.CertificateAuthority, error) {
	resp, err := e.CreateEndpoint(ctx, CreateRequest{CertificateAuthority: ca})
	if err != nil {
		return nil, err
	}
	response := resp.(CreateResponse)
	return response.CertificateAuthority, nil
}

func (e Endpoints) Delete(ctx context.Context, name string) error {
	_, err := e.DeleteEndpoint(ctx, DeleteRequest{Name: name})
	if err != nil {
		return err
	}
	return nil
}

func (e Endpoints) Get(ctx context.Context, name string) (*api.CertificateAuthority, error) {
	resp, err := e.GetEndpoint(ctx, GetRequest{Name: name})
	if err != nil {
		return nil, err
	}
	response := resp.(GetResponse)
	return response.CertificateAuthority, nil
}

func (e Endpoints) List(ctx context.Context) ([]*api.CertificateAuthority, error) {
	resp, err := e.ListEndpoint(ctx, ListRequest{})
	if err != nil {
		return nil, err
	}
	response := resp.(ListResponse)
	return response.CertificateAuthorities, nil
}

func (e Endpoints) Update(ctx context.Context, ca *api.CertificateAuthority) (*api.CertificateAuthority, error) {
	resp, err := e.UpdateEndpoint(ctx, UpdateRequest{CertificateAuthority: ca})
	if err != nil {
		return nil, err
	}
	response := resp.(UpdateResponse)
	return response.CertificateAuthority, nil
}

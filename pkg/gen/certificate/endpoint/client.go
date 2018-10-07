// Code generated by protoc-gen-gotemplate. DO NOT EDIT.

package endpoint

import (
	"context"

	"powerssl.io/pkg/api"
)

func (e Endpoints) Create(ctx context.Context, certificate *api.Certificate) (*api.Certificate, error) {
	resp, err := e.CreateEndpoint(ctx, CreateRequest{
		Certificate: certificate,
	})
	if err != nil {
		return nil, err
	}
	response := resp.(CreateResponse)
	return response.Certificate, nil
}

func (e Endpoints) Delete(ctx context.Context, name string) error {
	_, err := e.DeleteEndpoint(ctx, DeleteRequest{
		Name: name,
	})
	if err != nil {
		return err
	}
	return nil
}

func (e Endpoints) Get(ctx context.Context, name string) (*api.Certificate, error) {
	resp, err := e.GetEndpoint(ctx, GetRequest{
		Name: name,
	})
	if err != nil {
		return nil, err
	}
	response := resp.(GetResponse)
	return response.Certificate, nil
}

func (e Endpoints) List(ctx context.Context) ([]*api.Certificate, error) {
	resp, err := e.ListEndpoint(ctx, ListRequest{})
	if err != nil {
		return nil, err
	}
	response := resp.(ListResponse)
	return response.Certificates, nil
}

func (e Endpoints) Update(ctx context.Context, certificate *api.Certificate) (*api.Certificate, error) {
	resp, err := e.UpdateEndpoint(ctx, UpdateRequest{
		Certificate: certificate,
	})
	if err != nil {
		return nil, err
	}
	response := resp.(UpdateResponse)
	return response.Certificate, nil
}

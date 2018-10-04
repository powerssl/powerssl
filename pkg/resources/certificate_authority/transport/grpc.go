package transport

import (
	"context"

	//"google.golang.org/grpc"

	oldcontext "golang.org/x/net/context"

	//"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	types "github.com/gogo/protobuf/types"

	"powerssl.io/pkg/api"
	apiv1 "powerssl.io/pkg/api/v1"
	"powerssl.io/pkg/resources/certificate_authority/endpoints"
	//"powerssl.io/pkg/resources/certificate_authority/service"
	"powerssl.io/pkg/resources/transport"
)

type grpcServer struct {
	create grpctransport.Handler
	delete grpctransport.Handler
	get    grpctransport.Handler
	list   grpctransport.Handler
	update grpctransport.Handler
}

// NewGRPCServer makes a set of endpoints available as a gRPC CertificateAuthorityServiceServer.
func NewGRPCServer(endpoints endpoints.Endpoints, logger log.Logger) apiv1.CertificateAuthorityServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}

	return &grpcServer{
		create: grpctransport.NewServer(
			endpoints.CreateEndpoint,
			decodeGRPCCreateRequest,
			encodeGRPCCreateResponse,
			options...,
		),
		delete: grpctransport.NewServer(
			endpoints.DeleteEndpoint,
			decodeGRPCDeleteRequest,
			encodeGRPCDeleteResponse,
			options...,
		),
		get: grpctransport.NewServer(
			endpoints.GetEndpoint,
			decodeGRPCGetRequest,
			encodeGRPCGetResponse,
			options...,
		),
		list: grpctransport.NewServer(
			endpoints.ListEndpoint,
			decodeGRPCListRequest,
			encodeGRPCListResponse,
			options...,
		),
		update: grpctransport.NewServer(
			endpoints.UpdateEndpoint,
			decodeGRPCUpdateRequest,
			encodeGRPCUpdateResponse,
			options...,
		),
	}
}

func (s *grpcServer) CreateCertificateAuthority(ctx oldcontext.Context, req *apiv1.CreateCertificateAuthorityRequest) (*apiv1.CertificateAuthority, error) {
	_, rep, err := s.create.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.CertificateAuthority), nil
}

func (s *grpcServer) DeleteCertificateAuthority(ctx oldcontext.Context, req *apiv1.DeleteCertificateAuthorityRequest) (*types.Empty, error) {
	_, rep, err := s.delete.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*types.Empty), nil
}

func (s *grpcServer) GetCertificateAuthority(ctx oldcontext.Context, req *apiv1.GetCertificateAuthorityRequest) (*apiv1.CertificateAuthority, error) {
	_, rep, err := s.get.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.CertificateAuthority), nil
}

func (s *grpcServer) ListCertificateAuthorities(ctx oldcontext.Context, req *apiv1.ListCertificateAuthoritiesRequest) (*apiv1.ListCertificateAuthoritiesResponse, error) {
	_, rep, err := s.list.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.ListCertificateAuthoritiesResponse), nil
}

func (s *grpcServer) UpdateCertificateAuthority(ctx oldcontext.Context, req *apiv1.UpdateCertificateAuthorityRequest) (*apiv1.CertificateAuthority, error) {
	_, rep, err := s.update.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*apiv1.CertificateAuthority), nil
}

// NewGRPCClient returns an AddService backed by a gRPC server at the other end
// of the conn. The caller is responsible for constructing the conn, and
// eventually closing the underlying transport. We bake-in certain middlewares,
// implementing the client library pattern.
// func NewGRPCClient(conn *grpc.ClientConn, logger log.Logger) service.Service {
// 	// global client middlewares
// 	options := []grpctransport.ClientOption{}
//
// 	var createEndpoint endpoint.Endpoint
// 	{
// 		createEndpoint = grpctransport.NewClient(
// 			conn,
// 			"apiv1.CreateCertificateAuthority",
// 			"Create",
// 			encodeGRPCCreateRequest,
// 			decodeGRPCCreateResponse,
// 			apiv1.CertificateAuthority{},
// 			options...,
// 		).Endpoint()
// 	}
//
// 	var deleteEndpoint endpoint.Endpoint
// 	{
// 		deleteEndpoint = grpctransport.NewClient(
// 			conn,
// 			"apiv1.DeleteCertificateAuthority",
// 			"Delete",
// 			encodeGRPCDeleteRequest,
// 			decodeGRPCDeleteResponse,
// 			types.Empty{},
// 			options...,
// 		).Endpoint()
// 	}
//
// 	var getEndpoint endpoint.Endpoint
// 	{
// 		getEndpoint = grpctransport.NewClient(
// 			conn,
// 			"apiv1.GetCertificateAuthority",
// 			"Get",
// 			encodeGRPCGetRequest,
// 			decodeGRPCGetResponse,
// 			apiv1.CertificateAuthority{},
// 			options...,
// 		).Endpoint()
// 	}
//
// 	var listEndpoint endpoint.Endpoint
// 	{
// 		getEndpoint = grpctransport.NewClient(
// 			conn,
// 			"apiv1.ListCertificateAuthorities",
// 			"Get",
// 			encodeGRPCListRequest,
// 			decodeGRPCListResponse,
// 			apiv1.ListCertificateAuthoritiesResponse{},
// 			options...,
// 		).Endpoint()
// 	}
//
// 	var updateEndpoint endpoint.Endpoint
// 	{
// 		getEndpoint = grpctransport.NewClient(
// 			conn,
// 			"apiv1.UpdateCertificateAuthority",
// 			"Update",
// 			encodeGRPCUpdateRequest,
// 			decodeGRPCUpdateResponse,
// 			apiv1.CertificateAuthority{},
// 			options...,
// 		).Endpoint()
// 	}
//
// 	// Returning the endpoint.Set as a service.Service relies on the
// 	// endpoint.Set implementing the Service methods. That's just a simple bit
// 	// of glue code.
// 	return endpoints.Endpoints{
// 		CreateEndpoint: createEndpoint,
// 		DeleteEndpoint: deleteEndpoint,
// 		GetEndpoint:    getEndpoint,
// 		ListEndpoint:   listEndpoint,
// 		UpdateEndpoint: updateEndpoint,
// 	}
// }

func decodeGRPCCertificateAuthorities(certificateAuthorities []*apiv1.CertificateAuthority) []api.CertificateAuthority {
	items := make([]api.CertificateAuthority, len(certificateAuthorities))
	for i, certificateAuthority := range certificateAuthorities {
		items[i] = decodeGRPCCertificateAuthority(certificateAuthority)
	}
	return items
}

func decodeGRPCCertificateAuthority(certificateAuthority *apiv1.CertificateAuthority) api.CertificateAuthority {
	return api.CertificateAuthority{
		TypeMeta:   transport.DecodeGRPCTypeMeta(certificateAuthority.GetTypeMeta()),
		ObjectMeta: transport.DecodeGRPCObjectMeta(certificateAuthority.GetObjectMeta()),
		Spec: api.CertificateAuthoritySpec{
			Vendor: certificateAuthority.GetSpec().GetVendor(),
		},
	}
}

func encodeGRPCCertificateAuthorities(certificateAuthorities []api.CertificateAuthority) []*apiv1.CertificateAuthority {
	items := make([]*apiv1.CertificateAuthority, len(certificateAuthorities))
	for i, certificateAuthority := range certificateAuthorities {
		items[i] = encodeGRPCCertificateAuthority(certificateAuthority)
	}
	return items
}

func encodeGRPCCertificateAuthority(certificateAuthority api.CertificateAuthority) *apiv1.CertificateAuthority {
	return &apiv1.CertificateAuthority{
		TypeMeta:   transport.EncodeGRPCTypeMeta(certificateAuthority.TypeMeta),
		ObjectMeta: transport.EncodeGRPCObjectMeta(certificateAuthority.ObjectMeta),
		Spec: &apiv1.CertificateAuthoritySpec{
			Vendor: certificateAuthority.Spec.Vendor,
		},
	}
}

func decodeGRPCCreateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.CreateCertificateAuthorityRequest)
	return endpoints.CreateRequest{
		CertificateAuthority: decodeGRPCCertificateAuthority(req.GetCertificateAuthority()),
	}, nil
}

func decodeGRPCCreateResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.CertificateAuthority)
	return endpoints.CreateResponse{
		CertificateAuthority: decodeGRPCCertificateAuthority(reply),
	}, nil
}

func decodeGRPCDeleteRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.DeleteCertificateAuthorityRequest)
	return endpoints.DeleteRequest{
		Name: req.GetName(),
	}, nil
}

func decodeGRPCDeleteResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	return endpoints.DeleteResponse{}, nil
}

func decodeGRPCGetRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.GetCertificateAuthorityRequest)
	return endpoints.GetRequest{
		Name: req.GetName(),
	}, nil
}

func decodeGRPCGetResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.CertificateAuthority)
	return endpoints.GetResponse{
		CertificateAuthority: decodeGRPCCertificateAuthority(reply),
	}, nil
}

func decodeGRPCListRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return endpoints.ListRequest{}, nil
}

func decodeGRPCListResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.ListCertificateAuthoritiesResponse)
	return endpoints.ListResponse{
		CertificateAuthorities: decodeGRPCCertificateAuthorities(reply.GetItems()),
	}, nil
}

func decodeGRPCUpdateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*apiv1.UpdateCertificateAuthorityRequest)
	return endpoints.UpdateRequest{
		CertificateAuthority: decodeGRPCCertificateAuthority(req.GetCertificateAuthority()),
	}, nil
}

func decodeGRPCUpdateResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*apiv1.CertificateAuthority)
	return endpoints.UpdateResponse{
		CertificateAuthority: decodeGRPCCertificateAuthority(reply),
	}, nil
}

func encodeGRPCCreateResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.CreateResponse)
	return encodeGRPCCertificateAuthority(resp.CertificateAuthority), nil
}

func encodeGRPCCreateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoints.CreateRequest)
	return &apiv1.CreateCertificateAuthorityRequest{
		CertificateAuthority: encodeGRPCCertificateAuthority(req.CertificateAuthority),
	}, nil
}

func encodeGRPCDeleteResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func encodeGRPCDeleteRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoints.DeleteRequest)
	return &apiv1.DeleteCertificateAuthorityRequest{
		Name: req.Name,
	}, nil
}

func encodeGRPCGetResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.GetResponse)
	return encodeGRPCCertificateAuthority(resp.CertificateAuthority), nil
}

func encodeGRPCGetRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoints.GetRequest)
	return &apiv1.GetCertificateAuthorityRequest{
		Name: req.Name,
	}, nil
}

func encodeGRPCListResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.ListResponse)
	return encodeGRPCCertificateAuthorities(resp.CertificateAuthorities), nil
}

func encodeGRPCListRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &apiv1.ListCertificateAuthoritiesRequest{}, nil
}

func encodeGRPCUpdateResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.UpdateResponse)
	return encodeGRPCCertificateAuthority(resp.CertificateAuthority), nil
}

func encodeGRPCUpdateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoints.UpdateRequest)
	return &apiv1.UpdateCertificateAuthorityRequest{
		CertificateAuthority: encodeGRPCCertificateAuthority(req.CertificateAuthority),
	}, nil
}

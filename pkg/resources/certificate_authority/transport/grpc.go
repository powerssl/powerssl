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
	pb "powerssl.io/pkg/api/v1"
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
func NewGRPCServer(endpoints endpoints.Endpoints, logger log.Logger) pb.CertificateAuthorityServiceServer {
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

func (s *grpcServer) CreateCertificateAuthority(ctx oldcontext.Context, req *pb.CreateCertificateAuthorityRequest) (*pb.CertificateAuthority, error) {
	_, rep, err := s.create.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CertificateAuthority), nil
}

func (s *grpcServer) DeleteCertificateAuthority(ctx oldcontext.Context, req *pb.DeleteCertificateAuthorityRequest) (*types.Empty, error) {
	_, rep, err := s.delete.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*types.Empty), nil
}

func (s *grpcServer) GetCertificateAuthority(ctx oldcontext.Context, req *pb.GetCertificateAuthorityRequest) (*pb.CertificateAuthority, error) {
	_, rep, err := s.get.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CertificateAuthority), nil
}

func (s *grpcServer) ListCertificateAuthorities(ctx oldcontext.Context, req *pb.ListCertificateAuthoritiesRequest) (*pb.ListCertificateAuthoritiesResponse, error) {
	_, rep, err := s.get.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ListCertificateAuthoritiesResponse), nil
}

func (s *grpcServer) UpdateCertificateAuthority(ctx oldcontext.Context, req *pb.UpdateCertificateAuthorityRequest) (*pb.CertificateAuthority, error) {
	_, rep, err := s.update.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CertificateAuthority), nil
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
// 			"pb.CreateCertificateAuthority",
// 			"Create",
// 			encodeGRPCCreateRequest,
// 			decodeGRPCCreateResponse,
// 			pb.CertificateAuthority{},
// 			options...,
// 		).Endpoint()
// 	}
//
// 	var deleteEndpoint endpoint.Endpoint
// 	{
// 		deleteEndpoint = grpctransport.NewClient(
// 			conn,
// 			"pb.DeleteCertificateAuthority",
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
// 			"pb.GetCertificateAuthority",
// 			"Get",
// 			encodeGRPCGetRequest,
// 			decodeGRPCGetResponse,
// 			pb.CertificateAuthority{},
// 			options...,
// 		).Endpoint()
// 	}
//
// 	var listEndpoint endpoint.Endpoint
// 	{
// 		getEndpoint = grpctransport.NewClient(
// 			conn,
// 			"pb.ListCertificateAuthorities",
// 			"Get",
// 			encodeGRPCListRequest,
// 			decodeGRPCListResponse,
// 			pb.ListCertificateAuthoritiesResponse{},
// 			options...,
// 		).Endpoint()
// 	}
//
// 	var updateEndpoint endpoint.Endpoint
// 	{
// 		getEndpoint = grpctransport.NewClient(
// 			conn,
// 			"pb.UpdateCertificateAuthority",
// 			"Update",
// 			encodeGRPCUpdateRequest,
// 			decodeGRPCUpdateResponse,
// 			pb.CertificateAuthority{},
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

func decodeGRPCCertificateAuthority(certificateAuthority *pb.CertificateAuthority) api.CertificateAuthority {
	return api.CertificateAuthority{
		TypeMeta:   transport.DecodeGRPCTypeMeta(certificateAuthority.TypeMeta),
		ObjectMeta: transport.DecodeGRPCObjectMeta(certificateAuthority.ObjectMeta),
		Spec: api.CertificateAuthoritySpec{
			Vendor: certificateAuthority.Spec.Vendor,
		},
	}
}

func encodeGRPCCertificateAuthority(certificateAuthority api.CertificateAuthority) *pb.CertificateAuthority {
	return &pb.CertificateAuthority{
		TypeMeta:   transport.EncodeGRPCTypeMeta(certificateAuthority.TypeMeta),
		ObjectMeta: transport.EncodeGRPCObjectMeta(certificateAuthority.ObjectMeta),
		Spec: &pb.CertificateAuthoritySpec{
			Vendor: certificateAuthority.Spec.Vendor,
		},
	}
}

func decodeGRPCCreateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CreateCertificateAuthorityRequest)
	return endpoints.CreateRequest{
		CertificateAuthority: decodeGRPCCertificateAuthority(req.CertificateAuthority),
	}, nil
}

func decodeGRPCCreateResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.CertificateAuthority)
	return endpoints.CreateResponse{
		CertificateAuthority: decodeGRPCCertificateAuthority(reply),
	}, nil
}

func decodeGRPCDeleteRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.DeleteCertificateAuthorityRequest)
	return endpoints.DeleteRequest{
		Name: req.Name,
	}, nil
}

func decodeGRPCDeleteResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	return endpoints.DeleteResponse{}, nil
}

func decodeGRPCGetRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	// TODO
	return nil, nil
}

func decodeGRPCGetResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	// TODO
	return nil, nil
}

func decodeGRPCListRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	// TODO
	return nil, nil
}

func decodeGRPCListResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	// TODO
	return nil, nil
}

func decodeGRPCUpdateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	// TODO
	return nil, nil
}

func decodeGRPCUpdateResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	// TODO
	return nil, nil
}

func encodeGRPCCreateResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.CreateResponse)
	return encodeGRPCCertificateAuthority(resp.CertificateAuthority), nil
}

func encodeGRPCCreateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoints.CreateRequest)
	return &pb.CreateCertificateAuthorityRequest{
		CertificateAuthority: encodeGRPCCertificateAuthority(req.CertificateAuthority),
	}, nil
}

func encodeGRPCDeleteResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func encodeGRPCDeleteRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoints.DeleteRequest)
	return &pb.DeleteCertificateAuthorityRequest{
		Name: req.Name,
	}, nil
}

func encodeGRPCGetResponse(_ context.Context, response interface{}) (interface{}, error) {
	// TODO
	return nil, nil
}

func encodeGRPCGetRequest(_ context.Context, request interface{}) (interface{}, error) {
	// TODO
	return nil, nil
}

func encodeGRPCListResponse(_ context.Context, response interface{}) (interface{}, error) {
	// TODO
	return nil, nil
}

func encodeGRPCListRequest(_ context.Context, request interface{}) (interface{}, error) {
	// TODO
	return nil, nil
}

func encodeGRPCUpdateResponse(_ context.Context, response interface{}) (interface{}, error) {
	// TODO
	return nil, nil
}

func encodeGRPCUpdateRequest(_ context.Context, request interface{}) (interface{}, error) {
	// TODO
	return nil, nil
}
